package main

import (
	"log"

	"github.com/BrianToro/ISO8583/infrastructure/database"
	"github.com/BrianToro/ISO8583/interfaces"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewDatabaseFactory(database.InstanceMongoDB)
	if err != nil {
		log.Println(err)
	}

	s := interfaces.NewAPIServer(":8081", &db)
	err = s.Run()
	if err != nil {
		log.Println(err)
	}
}
