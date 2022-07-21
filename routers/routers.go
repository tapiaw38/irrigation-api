package routers

import (
	"github.com/gorilla/mux"
)

// UserRoutes is a function that returns a router for the user routes
func (ur *UserRouter) UserRoutes() *mux.Router {

	r := mux.NewRouter()
	//router := r.PathPrefix("/").Subrouter()
	r.HandleFunc("/register", ur.CreateUserHandler).Methods("POST")
	r.HandleFunc("/login", ur.LoginHandler).Methods("POST")
	r.HandleFunc("/all", ur.GetUsersHandler).Methods("GET")
	r.HandleFunc("/{id}", ur.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/username/{username}", ur.GetUserByUsernameHandler).Methods("GET")
	r.HandleFunc("/update/{id}", ur.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/partial/{id}", ur.PartialUpdateUserHandler).Methods("PUT")
	r.HandleFunc("/update/avatar/{id}", ur.UploadAvatarHandler).Methods("PUT")

	return r
}

// ProducerRoutes is a function that returns a router for the product routes
func (pr *ProducerRouter) ProducerRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", pr.CreateProducersHandler).Methods("POST")
	r.HandleFunc("/all", pr.GetProducersHandler).Methods("GET")
	r.HandleFunc("/{id}", pr.GetProducerByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", pr.UpdateProducerHandler).Methods("PUT")
	r.HandleFunc("/partial/{id}", pr.PartialUpdateProducerHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", pr.DeleteProducerHandler).Methods("DELETE")

	return r
}

// ProductionRoutes is a function that returns a router for the production routes
func (pr *ProductionRouter) ProductionRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", pr.CreateProductionHandler).Methods("POST")
	r.HandleFunc("/all", pr.GetProductionsHandler).Methods("GET")
	r.HandleFunc("/{id}", pr.GetProductionByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", pr.UpdateProductionHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", pr.DeleteProductionHandler).Methods("DELETE")
	r.HandleFunc("/upload/picture/{id}", pr.UploadPictureHandler).Methods("PUT")

	return r
}

// SectionRoutes is a function that returns a router for the section routes
func (sr *SectionRouter) SectionRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", sr.CreateSectionsHandler).Methods("POST")
	r.HandleFunc("/all", sr.GetSectionsHandler).Methods("GET")
	r.HandleFunc("/{id}", sr.GetSectionByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", sr.UpdateSectionHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", sr.DeleteSectionHandler).Methods("DELETE")

	return r
}

// IntakeRoutes is a function that returns a router for the intake routes
func (ir *IntakeRouter) IntakeRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", ir.CreateIntakesHandler).Methods("POST")
	r.HandleFunc("/all", ir.GetIntakesHandler).Methods("GET")
	r.HandleFunc("/{id}", ir.GetIntakeByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", ir.UpdateIntakeHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", ir.DeleteIntakeHandler).Methods("DELETE")
	r.HandleFunc("/production/{id}", ir.CreateIntakeProductionHandler).Methods("POST")
	r.HandleFunc("/production/update/{id}", ir.UpdateIntakeProductionHandler).Methods("PUT")
	r.HandleFunc("/production/delete/{id}", ir.DeleteIntakeProductionHandler).Methods("POST")

	return r
}

// TurnRoutes is a function that returns a router for the turn routes
func (tr *TurnRouter) TurnRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", tr.CreateTurnHandler).Methods("POST")
	// r.HandleFunc("/all", tr.GetTurnsHandler).Methods("GET")
	// r.HandleFunc("/{id}", tr.GetTurnByIDHandler).Methods("GET")
	// r.HandleFunc("/update/{id}", tr.UpdateTurnHandler).Methods("PUT")
	// r.HandleFunc("/delete/{id}", tr.DeleteTurnHandler).Methods("DELETE")
	// r.HandleFunc("/intake/{id}", tr.CreateTurnIntakeHandler).Methods("POST")
	// r.HandleFunc("/intake/delete/{id}", tr.DeleteTurnIntakeHandler).Methods("POST")

	return r
}
