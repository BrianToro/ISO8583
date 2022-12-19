package repositories

import "github.com/BrianToro/ISO8583/domain/models"

type TransactionsRepository interface {
	Create(t *models.Transaction) error
	Get(id string) (*models.Transaction, error)
}
