package server

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/irrigation-api/cache"
	"github.com/tapiaw38/irrigation-api/database"
	"github.com/tapiaw38/irrigation-api/middlewares"
	"github.com/tapiaw38/irrigation-api/repository"
	"github.com/tapiaw38/irrigation-api/utils"
	"github.com/tapiaw38/irrigation-api/websocket"
)

var (
	once sync.Once
)

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
	S3() *utils.S3Client
	Redis() *cache.RedisCache
}

type Config struct {
	DatabaseURL        string
	Port               string
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSBucket          string
	RedisHost          string
	RedisDB            int
	RedisExpires       time.Duration
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
	s3     *utils.S3Client
	redis  *cache.RedisCache
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func (b *Broker) S3() *utils.S3Client {
	return b.s3
}

func (b *Broker) Redis() *cache.RedisCache {
	return b.redis
}

func NewServer(config *Config) (*Broker, error) {
	if config.DatabaseURL == "" {
		return nil, errors.New("database url is required")
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
		s3: utils.NewSession(&utils.S3Config{
			AWSRegion:          config.AWSRegion,
			AWSAccessKeyID:     config.AWSAccessKeyID,
			AWSSecretAccessKey: config.AWSSecretAccessKey,
			AWSBucket:          config.AWSBucket,
		}),
		redis: cache.NewRedisCache(&cache.RedisCache{
			Host:    config.RedisHost,
			DB:      config.RedisDB,
			Expires: config.RedisExpires,
		}),
	}
	return broker, nil
}

func (b *Broker) Serve(binder func(s Server, r *mux.Router)) {

	// start router mux
	b.router = mux.NewRouter()
	binder(b, b.router)

	// connecting database
	once.Do(func() {
		db, err := database.ConnectPostgres(b.Config().DatabaseURL)

		if err != nil {
			panic(err)
		}

		err = db.MakeMigration(b.Config().DatabaseURL)

		if err != nil {
			panic(err)
		}

		log.Println("Migrate database")

		repository.SetRepository(db)
	})

	// run websocket hub
	go b.hub.Run()

	// Mount the middleware
	b.router.Use(middlewares.MiddlewareLog)
	//router.Use(middlewares.MiddlewareAuth)
	handler := cors.AllowAll().Handler(b.router)

	log.Println("listening on port", b.config.Port)

	if err := http.ListenAndServe(":"+b.config.Port, handler); err != nil {
		log.Fatal(err)
	}
}
