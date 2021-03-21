package models

type Subject struct {
	//Any string `json:",any"`
	//Street2        string `json:"street2"`

	Id             int    `json:"id,omitempty"`
	CustomId       string `json:"custom_id,omitempty"`
	Name           string `json:"name"`
	Street         string `json:"street"`
	City           string `json:"city"`
	Zip            string `json:"zip"`
	Country        string `json:"country,omitempty"`
	RegistrationNo string `json:"registration_no,omitempty"`
	VatNo          string `json:"vat_no,omitempty"`
	BankAccount    string `json:"bank_account,omitempty"`
	Iban           string `json:"iban,omitempty"`
	VariableSymbol string `json:"variable_symbol,omitempty"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	EmailCopy      string `json:"email_copy,omitempty"`
	Phone          string `json:"phone"`
	Web            string `json:"web,omitempty"`
}
