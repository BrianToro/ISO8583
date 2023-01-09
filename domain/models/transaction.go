package models

import (
	"errors"
	"time"
)

type Transaction struct {
	Id          string    `json:"id"`
	Amount      int       `json:"amount"`
	Pan         string    `json:"pan"`
	Gateway     string    `json:"gateway"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}

func NewTransaction(
	id string,
	amount int,
	pan string,
	gateway string,
	status string,
	description string,
	created time.Time,
) (*Transaction, error) {

	if amount < 500 {
		return nil, errors.New("transaction amount should be more than 500")
	}

	return &Transaction{
		Id:          id,
		Amount:      amount,
		Pan:         pan,
		Gateway:     gateway,
		Status:      status,
		Description: description,
		Created:     created,
	}, nil
}
