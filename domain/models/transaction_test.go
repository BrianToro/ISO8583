package models_test

import (
	"testing"
	"time"

	"github.com/BrianToro/ISO8583/domain/models"
)

func TestTransactionModel(t *testing.T) {
	var (
		amount      = 500
		id          = "test"
		pan         = "test"
		gateway     = "test"
		status      = "approved"
		description = "test"
		created     = time.Now()
	)
	transaction, err := models.NewTransaction(
		id,
		amount,
		pan,
		gateway,
		status,
		description,
		created,
	)

	if err != nil {
		t.Error(err)
	}

	if transaction.Amount != amount {
		t.Errorf("Expected: %v, got: %v", amount, transaction.Amount)
	}

	if transaction.Id != id {
		t.Errorf("Expected: %v, got: %v", id, transaction.Id)
	}

	if transaction.Pan != pan {
		t.Errorf("Expected: %v, got: %v", pan, transaction.Pan)
	}

	if transaction.Gateway != gateway {
		t.Errorf("Expected: %v, got: %v", gateway, transaction.Gateway)
	}

	if transaction.Status != status {
		t.Errorf("Expected: %v, got: %v", status, transaction.Status)
	}

	if transaction.Description != description {
		t.Errorf("Expected: %v, got: %v", description, transaction.Description)
	}

	if transaction.Created != created {
		t.Errorf("Expected: %v, got: %v", created, transaction.Created)
	}
}
