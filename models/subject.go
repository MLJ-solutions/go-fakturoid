package models

import "time"

type Subject struct {
	//Street2           string `json:"street2"` @deprecated
	Id               int       `json:"id,omitempty"`
	Custom_id        string    `json:"custom_id,omitempty"`
	Type             string    `json:"type"`
	Name             string    `json:"name"`
	Street           string    `json:"street"`
	City             string    `json:"city,omitempty"`
	Zip              string    `json:"zip,omitempty"`
	Country          string    `json:"country,omitempty"`
	RegistrationNo   string    `json:"registration_no,omitempty"`
	VatNo            string    `json:"vat_no,omitempty"`
	LocalVatNo       string    `json:"local_vat_no,omitempty"`
	BankAccount      string    `json:"bank_account"`
	Iban             string    `json:"iban"`
	VariableSymbol   string    `json:"variable_symbol,omitempty"`
	EnabledReminders bool      `json:"enabled_reminders"`
	FullName         string    `json:"full_name,omitempty"`
	Email            string    `json:"email,omitempty"`
	EmailCopy        string    `json:"email_copy,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	Web              string    `json:"web,omitempty"`
	PrivateNote      string    `json:"private_note,omitempty"`
	AvatarUrl        string    `json:"avatar_url,omitempty"`
	HtmlUrl          string    `json:"html_url,omitempty"`
	Url              string    `json:"url,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
