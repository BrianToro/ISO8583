package application

import (
	"context"
	"time"

	"github.com/BrianToro/ISO8583/domain/models"
	"github.com/BrianToro/ISO8583/domain/repositories"
	"github.com/google/uuid"
)

type (
	CreateTransactionUseCase interface {
		Execute(context.Context, *models.Transaction) error
	}

	CreateTransactionIterator struct {
		transactionRepository repositories.TransactionsRepository
		ctxTimeout            time.Duration
	}
)

const (
	Pending  string = "pending"
	Approved string = "approved"
	Failed   string = "failed"
)

func NewCreateTransactionInteractor(
	transactionRepository repositories.TransactionsRepository,
	t time.Duration,
) *CreateTransactionIterator {
	return &CreateTransactionIterator{
		transactionRepository: transactionRepository,
		ctxTimeout:            t,
	}
}

func (us *CreateTransactionIterator) Execute(
	ctx context.Context,
	createTransaction *models.CreateTransaction,
) error {
	transaction, err := models.NewTransaction(
		uuid.New().String(),
		createTransaction.Amount,
		createTransaction.Pan,
		"testing_gateway",
		Pending,
		createTransaction.Description,
		time.Now(),
	)
	if err != nil {
		return err
	}

	err = us.transactionRepository.Create(transaction)
	if err != nil {
		return err
	}

	return nil
}
