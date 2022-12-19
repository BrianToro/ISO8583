package models

type Transaction struct {
	Id     string `json:"id"`
	Amount int    `json:"amount"`
	Pan    string `json:"pan"`
}

func NewTransaction(id string, amount int, pan string) *Transaction {
	return &Transaction{
		Id:     id,
		Amount: amount,
		Pan:    pan,
	}
}

func (t *Transaction) GetId() string {
	return t.Id
}

func (t *Transaction) GetAmount() int {
	return t.Amount
}

func (t *Transaction) GetPan() string {
	return t.Pan
}
