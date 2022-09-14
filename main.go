package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tapiaw38/irrigation-api/routers"
	"github.com/tapiaw38/irrigation-api/server"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Println(err, "Error loading .env file")
	}

	// initialize new server
	DATABASE_URL := os.Getenv("DATABASE_URL")
	PORT := os.Getenv("PORT")
	// create an AWS session which can be
	AWS_REGION := os.Getenv("AWS_REGION")
	AWS_ACCESS_KEY_ID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY := os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWS_BUCKET := os.Getenv("AWS_BUCKET")
	// connect Redis
	REDIS_HOST := os.Getenv("REDIS_HOST")

	s, err := server.NewServer(&server.Config{
		DatabaseURL:        DATABASE_URL,
		Port:               PORT,
		AWSRegion:          AWS_REGION,
		AWSAccessKeyID:     AWS_ACCESS_KEY_ID,
		AWSSecretAccessKey: AWS_SECRET_ACCESS_KEY,
		AWSBucket:          AWS_BUCKET,
		RedisHost:          REDIS_HOST,
		RedisDB:            0,
		RedisExpires:       10,
	})

	if err != nil {
		log.Fatal(err)
	}

	// start server
	s.Serve(routers.BinderRoutes)
}
