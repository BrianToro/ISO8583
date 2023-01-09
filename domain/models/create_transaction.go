package models

type CreateTransaction struct {
	Amount      int    `json:"amount" validate:"required,number"`
	Iva         int    `json:"iva" validate:"number"`
	Inc         int    `json:"inc" validate:"number"`
	Pan         string `json:"pan" validate:"required,numeric"`
	Reference   string `json:"reference" validate:"required"`
	DeviceId    string `json:"deviceId" validate:"required,uuid"`
	Description string `json:"description" validate:"required,alphanum"`
}
