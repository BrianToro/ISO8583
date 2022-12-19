package interfaces

import (
	"github.com/BrianToro/ISO8583/application"
	"github.com/gofiber/fiber/v2"
)

type APIServer struct {
	Port string
	
}

func NewAPIServer(port string) *APIServer {
	return &APIServer{
		Port: port,
	}
}

func (s *APIServer) Run() error {
	app := fiber.New()
	api := app.Group("/api")

	// Register the health route
	api.Get("/", s.health)

	// Register transaction routes
	api.Post("/transaction", s.createTransaction)

	return app.Listen(s.Port)
}

func (s *APIServer) health(c *fiber.Ctx) error {
	return c.SendString("API working!")
}

func (s *APIServer) createTransaction(c *fiber.Ctx) error {
	act := application.NewCreateTransactionInteractor()

	return c.
}
