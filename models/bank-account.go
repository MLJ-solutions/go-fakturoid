package models

type BankAccount struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	Number            string `json:"number"`
	Iban              string `json:"iban"`
	SwiftBic          string `json:"swift_bic"`
	Pairing           bool   `json:"pairing"`
	PaymentAdjustment bool   `json:"payment_adjustment"`
}
