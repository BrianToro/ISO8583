package models

type Transaction struct {
	id     string
	amount int
	pan    string
}

func NewTransaction(id string, amount int, pan string) *Transaction {
	return &Transaction{
		id:     id,
		amount: amount,
		pan:    pan,
	}
}

func (t *Transaction) GetId() string {
	return t.id
}

func (t *Transaction) GetAmount() int {
	return t.amount
}

func (t *Transaction) GetPan() string {
	return t.pan
}
