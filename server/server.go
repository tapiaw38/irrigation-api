package server

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/irrigation-api/database"
	"github.com/tapiaw38/irrigation-api/middlewares"
	"github.com/tapiaw38/irrigation-api/repository"
)

var (
	once sync.Once
)

type Server interface {
	Config() *Config
}

type Config struct {
	DatabaseURL string
	Port        string
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(config *Config) (*Broker, error) {
	if config.DatabaseURL == "" {
		return nil, errors.New("database url is required")
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	b := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return b, nil
}

func (b *Broker) Serve(binder func(r *mux.Router)) {

	// init database
	once.Do(func() {
		db, err := database.ConnectPostgres(b.Config().DatabaseURL)

		if err != nil {
			panic(err)
		}

		err = db.MakeMigration()

		if err != nil {
			panic(err)
		}

		log.Println("Migrate database")

		repository.SetRepository(db)
	})

	// start router mux
	b.router = mux.NewRouter()
	binder(b.router)

	// Mount the middleware
	b.router.Use(middlewares.MiddlewareLog)
	//router.Use(middlewares.MiddlewareAuth)

	handler := cors.AllowAll().Handler(b.router)

	log.Println("listening on port", b.config.Port)

	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal(err)
	}
}
