package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/claim"
	"github.com/tapiaw38/irrigation-api/models"
	"github.com/tapiaw38/irrigation-api/repository"
	"github.com/tapiaw38/irrigation-api/server"
)

// Login is function to return the user
func Login(email string, password string) (models.User, bool) {

	u, find := repository.CheckUser(email)

	if !find {
		return u, false
	}

	if !u.PasswordMatch(password) {
		return u, false
	}

	return u, true
}

// LoginHandler handles the request to login a user
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "The email or password are invalid "+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "The user email is required", 400)
		return
	}

	foundUser, exist := Login(u.Email, u.Password)

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

	data := models.LoginResponse{
		User: models.User{
			ID:          foundUser.ID,
			FirstName:   foundUser.FirstName,
			LastName:    foundUser.LastName,
			Username:    foundUser.Username,
			Email:       foundUser.Email,
			Picture:     foundUser.Picture,
			PhoneNumber: foundUser.PhoneNumber,
			Address:     foundUser.Address,
			IsActive:    foundUser.IsActive,
			IsAdmin:     foundUser.IsAdmin,
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

// CreateUserHandler handles the request to create a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	user, err := repository.CreateUser(ctx, &u)

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
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	users, err := repository.GetUsers(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get users "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", users)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetUserByIdHandler handles the request to get a user by id
func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserById(ctx, id)

	if err != nil {
		http.Error(w, "No record found with that id "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)

}

// GetUserByUsernameHandler handles the request to get a user by username
func GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	username := mux.Vars(r)["username"]

	if username == "" {
		http.Error(w, "An error occurred, username is required", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByUsername(ctx, username)

	if err != nil {
		http.Error(w, "No record found with that username "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)

}

// UpdateUserHandler handles the request to update a user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User

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
	user, err := repository.UpdateUser(ctx, id, u)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a user in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)
}

// PartialUpdateUserHandler handles the request to update a user
func PartialUpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User

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
	user, err := repository.PartialUpdateUser(ctx, id, u)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a user in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteUserHandler handles the request to delete a user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err := repository.DeleteUser(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred while trying to delete a record from the database"+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", nil)
	ResponseWithJson(w, response, http.StatusOK)
}

// UploadAvatarHandler handles the request to upload an avatar
func UploadAvatarHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]

		if id == "" {
			http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
			return
		}

		maxSize := int64(1024 * 1024 * 5) // 5MB

		err := r.ParseMultipartForm(maxSize)

		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Image too large. Max Size: %v", maxSize)
			return
		}

		file, fileHeader, err := r.FormFile("picture")

		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Could not get uploaded file")
			return
		}

		defer file.Close()

		// reused if we're uploading many files
		fileName, err := s.S3().UploadFileToS3(file, fileHeader, id)

		if err != nil {
			fmt.Fprintf(w, "Could not upload file error"+err.Error(), 400)
			return
		}

		ctx := r.Context()

		fileUrl := s.S3().GenerateUrl(fileName)

		u, err := repository.GetUserById(ctx, id)

		if err != nil {
			fmt.Fprintf(w, "Could not get user from database"+err.Error(), 400)
			return
		}

		u.Picture = fileUrl

		user, err := repository.PartialUpdateUser(ctx, id, u)

		if err != nil {
			http.Error(w, "An error occurred when trying to update a user in database "+err.Error(), 400)
			return
		}

		response := NewResponse(Message, "ok", user)
		ResponseWithJson(w, response, http.StatusOK)

	}
}
