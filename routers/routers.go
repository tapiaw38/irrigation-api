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
	router.HandleFunc("/{id}", ur.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/username/{username}", ur.GetUserByUsernameHandler).Methods("GET")
	router.HandleFunc("/update/{id}", ur.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/partial/{id}", ur.PartialUpdateUserHandler).Methods("PUT")

	return r
}

// ProducerRoutes is a function that returns a router for the product routes
func (pr *ProducerRouter) ProducerRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", pr.CreateProducersHandler).Methods("POST")
	//r.HandleFunc("/all", pr.GetProductsHandler).Methods("GET")
	//r.HandleFunc("/{id}", pr.GetProductByIdHandler).Methods("GET")
	//r.HandleFunc("/update/{id}", pr.UpdateProductHandler).Methods("PUT")
	//r.HandleFunc("/partial/{id}", pr.PartialUpdateProductHandler).Methods("PUT")

	return r
}
