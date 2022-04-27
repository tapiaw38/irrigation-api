package routers

import (
	"github.com/gorilla/mux"
)

// UserRoutes is a function that returns a router for the user routes
func (ur *UserRouter) UserRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", ur.CreateUserHandler).Methods("POST")
	r.HandleFunc("/login", ur.LoginHandler).Methods("POST")
	router := r.PathPrefix("/").Subrouter()
	//router.Use(middlewares.MiddlewareAuth)
	router.HandleFunc("/all", ur.GetUsersHandler).Methods("GET")
	router.HandleFunc("/get_by_id/{id}", ur.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/update/{id}", ur.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/delete/{id}", ur.DeleteUserHandler).Methods("DELETE")

	return r
}
