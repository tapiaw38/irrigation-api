package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/irrigation-api/middlewares"
	"github.com/tapiaw38/irrigation-api/routers"
	"github.com/tapiaw38/irrigation-api/storage"
)

func HandlerServer() {
	router := mux.NewRouter()

	users := &routers.UserRouter{
		Storage: &storage.UserStorage{
			Data: storage.NewConnection(),
		},
	}

	producers := &routers.ProducerRouter{
		Storage: &storage.ProducerStorage{
			Data: storage.NewConnection(),
		},
	}

	productions := &routers.ProductionRouter{
		Storage: &storage.ProductionStorage{
			Data: storage.NewConnection(),
		},
	}

	sections := &routers.SectionRouter{
		Storage: &storage.SectionStorage{
			Data: storage.NewConnection(),
		},
	}

	intakes := &routers.IntakeRouter{
		Storage: &storage.IntakeStorage{
			Data: storage.NewConnection(),
		},
	}

	turns := &routers.TurnRouter{
		Storage: &storage.TurnStorage{
			Data: storage.NewConnection(),
		},
	}

	// Mount the routers
	mount(router, "/users", users.UserRoutes())
	mount(router, "/producers", producers.ProducerRoutes())
	mount(router, "/productions", productions.ProductionRoutes())
	mount(router, "/sections", sections.SectionRoutes())
	mount(router, "/intakes", intakes.IntakeRoutes())
	mount(router, "/turns", turns.TurnRoutes())

	// Mount the middleware
	router.Use(middlewares.MiddlewareLog)
	router.Use(middlewares.MiddlewareAuth)

	handler := cors.AllowAll().Handler(router)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: handler,
	}

	log.Println("Listening on port " + PORT)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

// mount is a helper function to mount a router to a path
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
