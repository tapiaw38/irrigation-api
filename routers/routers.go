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
	r.HandleFunc("/all", pr.GetProducersHandler).Methods("GET")
	r.HandleFunc("/update/{id}", pr.UpdateProducerHandler).Methods("PUT")
	r.HandleFunc("/partial/{id}", pr.PartialUpdateProducerHandler).Methods("PUT")

	return r
}

// ProductionRoutes is a function that returns a router for the production routes
func (pr *ProductionRouter) ProductionRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", pr.CreateProductionHandler).Methods("POST")
	r.HandleFunc("/all", pr.GetProductionsHandler).Methods("GET")
	r.HandleFunc("/update/{id}", pr.UpdateProductionHandler).Methods("PUT")

	return r
}
