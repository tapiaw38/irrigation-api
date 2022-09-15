package routers

import (
	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/server"
)

// UserRoutes is a function that returns a router for the user routes
func UserRoutes(s server.Server) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/register", CreateUserHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/all", GetUsersHandler).Methods("GET")
	r.HandleFunc("/{id}", GetUserByIdHandler(s)).Methods("GET")
	r.HandleFunc("/username/{username}", GetUserByUsernameHandler).Methods("GET")
	r.HandleFunc("/update/{id}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/partial/{id}", PartialUpdateUserHandler).Methods("PUT")
	r.HandleFunc("/update/avatar/{id}", UploadAvatarHandler(s)).Methods("PUT")

	return r
}

// ProducerRoutes is a function that returns a router for the product routes
func ProducerRoutes(s server.Server) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", CreateProducersHandler(s)).Methods("POST")
	r.HandleFunc("/all", GetProducersHandler(s)).Methods("GET")
	r.HandleFunc("/{id}", GetProducerByIDHandler(s)).Methods("GET")
	r.HandleFunc("/update/{id}", UpdateProducerHandler(s)).Methods("PUT")
	r.HandleFunc("/partial/{id}", PartialUpdateProducerHandler(s)).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteProducerHandler(s)).Methods("DELETE")

	return r
}

// ProductionRoutes is a function that returns a router for the production routes
func ProductionRoutes(s server.Server) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", CreateProductionHandler).Methods("POST")
	r.HandleFunc("/all", GetProductionsHandler).Methods("GET")
	r.HandleFunc("/{id}", GetProductionByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", UpdateProductionHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteProductionHandler).Methods("DELETE")
	r.HandleFunc("/upload/picture/{id}", UploadPictureHandler(s)).Methods("PUT")

	return r
}

// SectionRoutes is a function that returns a router for the section routes
func SectionRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", CreateSectionsHandler).Methods("POST")
	r.HandleFunc("/all", GetSectionsHandler).Methods("GET")
	r.HandleFunc("/{id}", GetSectionByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", UpdateSectionHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteSectionHandler).Methods("DELETE")

	return r
}

// IntakeRoutes is a function that returns a router for the intake routes
func IntakeRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", CreateIntakesHandler).Methods("POST")
	r.HandleFunc("/all", GetIntakesHandler).Methods("GET")
	r.HandleFunc("/{id}", GetIntakeByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", UpdateIntakeHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteIntakeHandler).Methods("DELETE")
	r.HandleFunc("/production/{id}", CreateIntakeProductionHandler).Methods("POST")
	r.HandleFunc("/production/update/{id}", UpdateIntakeProductionHandler).Methods("PUT")
	r.HandleFunc("/production/delete/{id}", DeleteIntakeProductionHandler).Methods("POST")

	return r
}

// TurnRoutes is a function that returns a router for the turn routes
func TurnRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/create", CreateTurnHandler).Methods("POST")
	r.HandleFunc("/all", GetTurnsHandler).Methods("GET")
	r.HandleFunc("/{id}", GetTurnByIDHandler).Methods("GET")
	r.HandleFunc("/update/{id}", UpdateTurnHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteTurnHandler).Methods("DELETE")
	r.HandleFunc("/production/{id}", CreateTurnProductionHandler).Methods("POST")
	r.HandleFunc("/production/delete/{id}", DeleteTurnProductionHandler).Methods("POST")

	return r
}

func WebSocketRoutes(s server.Server) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", s.Hub().HandleWebSocket)

	return r
}
