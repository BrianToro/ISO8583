package infrastructure

import (
	"errors"
	"time"

	"github.com/BrianToro/ISO8583/adapter/repository"
	"github.com/BrianToro/ISO8583/application"
	"github.com/BrianToro/ISO8583/domain/models"
	repositories_infra "github.com/BrianToro/ISO8583/infrastructure/repositories"
	"github.com/BrianToro/ISO8583/infrastructure/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type APIServer struct {
	Port       string
	Repository *repository.NoSQL
}

func NewAPIServer(port string, repository *repository.NoSQL) *APIServer {
	return &APIServer{
		Port:       port,
		Repository: repository,
	}
}

func (s *APIServer) Run() error {
	app := fiber.New()
	app.Use(logger.New())
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
	r := repositories_infra.NewTransactionRepositoryImpl(*s.Repository)
	act := application.NewCreateTransactionInteractor(*r, 10*time.Second)
	createTransaction := &models.CreateTransaction{}
	if err := c.BodyParser(createTransaction); err != nil {
		return errors.New("error parsing the transaction")
	}
	errs := validations.ValidateStruct(createTransaction)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}
	return act.Execute(c.Context(), createTransaction)
}
