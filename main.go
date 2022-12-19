package main

import (
	"log"

	"github.com/BrianToro/ISO8583/interfaces"
)

func main() {
	s := interfaces.NewAPIServer(":8081")
	err := s.Run()
	if err != nil {
		log.Println("error starting the server")
	}
}
