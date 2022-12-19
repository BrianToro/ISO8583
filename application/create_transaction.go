package application

import (
	"context"
	"errors"
	"time"

	"github.com/BrianToro/ISO8583/domain/models"
	"github.com/BrianToro/ISO8583/domain/repositories"
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
	transaction *models.Transaction,
) error {
	if transaction.GetAmount() <= 0 {
		return errors.New("transaction has a invalid amount")
	}

	err := us.transactionRepository.Create(transaction)
	if err != nil {
		return err
	}

	return nil
}
