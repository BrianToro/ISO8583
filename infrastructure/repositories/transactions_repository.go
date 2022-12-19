package repositories_infra

import (
	"context"
	"errors"

	"github.com/BrianToro/ISO8583/adapter/repository"
	"github.com/BrianToro/ISO8583/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	transactionsCollection = "transctions"
)

type TransactionsRepositoryImpl struct {
	databaseRepository repository.NoSQL
}

func NewTransactionRepositoryImpl(r repository.NoSQL) *TransactionsRepositoryImpl {
	return &TransactionsRepositoryImpl{
		databaseRepository: r,
	}
}

func (t TransactionsRepositoryImpl) Create(transaction *models.Transaction) error {
	return t.databaseRepository.Store(context.Background(), transactionsCollection, transaction)
}

func (t TransactionsRepositoryImpl) Get(id string) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	if err := t.databaseRepository.FindOne(context.Background(), transactionsCollection, bson.M{"id": id}, nil, transaction); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return &models.Transaction{}, errors.New("error not found")
		default:
			return &models.Transaction{}, err
		}
	}

	return transaction, nil
}
