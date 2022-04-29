package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/claim"
	"github.com/tapiaw38/irrigation-api/models/user"
)

// UserRouter is the router for the user api
type UserRouter struct {
	Storage user.Storage
}

// LoginHandler handles the request to login a user
func (ur *UserRouter) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "The email or password are invalid "+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "The user email is required", 400)
		return
	}

	foundUser, exist := ur.Login(u.Email, u.Password)

	if !exist {
		http.Error(w, "The email or password are invalid", 400)
		return
	}

	if !foundUser.IsActive {
		http.Error(w, "This user is not activated, please contact your administrator", 400)
		return
	}

	jwtKey, err := claim.GenerateJWT(foundUser)

	if err != nil {
		http.Error(w, "An error occurred while generating the token"+err.Error(), 400)
		return
	}

	data := user.LoginResponse{
		User: user.User{
			ID:        foundUser.ID,
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			Username:  foundUser.Username,
			Email:     foundUser.Email,
			Picture:   foundUser.Picture,
			Address:   foundUser.Address,
			IsActive:  foundUser.IsActive,
			IsAdmin:   foundUser.IsAdmin,
		},
		AccessToken: jwtKey,
	}

	response := NewResponse(Message, "ok", data)
	ResponseWithJson(w, response, http.StatusOK)

	//coquie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

// Login is function to return the user
func (ur *UserRouter) Login(email string, password string) (user.User, bool) {

	u, find := ur.Storage.CheckUser(email)

	if !find {
		return u, false
	}

	if !u.PasswordMatch(password) {
		return u, false
	}

	return u, true
}

// CreateUserHandler handles the request to create a new user
func (ur *UserRouter) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	user, err := ur.Storage.CreateUser(ctx, &u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user in database "+err.Error(), 400)
		return
	}

	u.PasswordHash = ""

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusCreated)

}

// GetUsersHandler handles the request to get all users
func (ur *UserRouter) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	users, err := ur.Storage.GetUsers(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get users "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", users)
	ResponseWithJson(w, response, http.StatusOK)

}

// GetUserByIdHandler handles the request to get a user by id
func (ur *UserRouter) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	user, err := ur.Storage.GetUserById(ctx, id)

	if err != nil {
		http.Error(w, "No record found with that id "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)

}

// GetUserByUsernameHandler handles the request to get a user by username
func (ur *UserRouter) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	username := r.URL.Query().Get("username")

	user, err := ur.Storage.GetUserByUsername(ctx, username)

	if err != nil {
		http.Error(w, "No record found with that username "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)

}

// UpdateUserHandler handles the request to update a user
func (ur UserRouter) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u user.User

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	user, err := ur.Storage.UpdateUser(ctx, id, u)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a user in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteUserHandler handles the request to delete a user
func (ur UserRouter) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err := ur.Storage.DeleteUser(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred while trying to delete a record from the database"+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", nil)
	ResponseWithJson(w, response, http.StatusOK)
}
