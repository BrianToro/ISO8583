package database

import (
	"errors"

	"github.com/BrianToro/ISO8583/adapter/repository"
)

var (
	errInvalidDatabaseInstance = errors.New("invalid nosql db instance")
)

const (
	InstanceMongoDB int = 0
)

func NewDatabaseFactory(instance int) (repository.NoSQL, error) {
	switch instance {
	case InstanceMongoDB:
		return NewHandlerDB(NewMongoDBConfig())
	default:
		return nil, errInvalidDatabaseInstance
	}
}
