package models

type Line struct {
	Name      string `json:"name"`
	Quantity  string `json:"quantity"`
	UnitName  string `json:"unit_name"`
	UnitPrice string `json:"unit_price,float"`
	VatRate   int    `json:"vat_rate"`
}

type Invoice struct {
	Number                string `json:"number,omitempty"`
	CustomId              string `json:"custom_id,omitempty"`
	OrderNumber           string `json:"order_number,omitempty"`
	VariableSymbol        string `json:"variable_symbol,omitempty"`
	SubjectId             int    `json:"subject_id"`
	Currency              string `json:"currency"`
	PaymentMethod         string `json:"payment_method"`
	Due                   int    `json:"due"`
	IssuedOn              string `json:"issued_on"`
	TaxableFulfillmentDue string `json:"taxable_fulfillment_due"`
	BankAccountId         int    `json:"bank_account_id"`
	Lines                 []Line `json:"lines"`
}
