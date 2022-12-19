package database

import (
	"os"
	"time"
)

type config struct {
	host     string
	database string
	user     string
	password string

	ctxTimeout time.Duration
}

func NewMongoDBConfig() *config {
	return &config{
		host:       os.Getenv("MONGO_HOST"),
		database:   os.Getenv("MONGODB_DB"),
		password:   os.Getenv("MONGODB_ROOT_PASSWORD"),
		user:       os.Getenv("MONGODB_ROOT_USER"),
		ctxTimeout: 60 * time.Second,
	}
}
