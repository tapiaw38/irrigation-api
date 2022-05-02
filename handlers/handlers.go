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

	// Mount the routers
	mount(router, "/users", users.UserRoutes())
	mount(router, "/producers", producers.ProducerRoutes())

	// Mount the middleware
	router.Use(middlewares.MiddlewareLog)

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
