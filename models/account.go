package models

import "time"

const (
	VatModePayer            = "vat_payer"
	VatModeNotPayer         = "not_vat_payer"
	VatModeIdentifiedPerson = "identified_person"
)

type Account struct {
	Subdomain                    string    `json:"subdomain"`
	Plan                         string    `json:"plan"`
	PlanPrice                    int       `json:"plan_price"`
	Email                        string    `json:"email"`
	InvoiceEmail                 string    `json:"invoice_email"`
	Phone                        string    `json:"phone"`
	Web                          string    `json:"web"`
	Name                         string    `json:"name"`
	FullName                     string    `json:"full_name"`
	RegistrationNo               string    `json:"registration_no"`
	VatNo                        string    `json:"vat_no"`
	VatMode                      string    `json:"vat_mode"`
	VatPriceMode                 string    `json:"vat_price_mode"`
	Street                       string    `json:"street"`
	Street2                      string    `json:"street2"`
	City                         string    `json:"city"`
	Zip                          string    `json:"zip"`
	Country                      string    `json:"country"`
	BankAccount                  string    `json:"bank_account"`
	Iban                         string    `json:"iban"`
	SwiftBic                     string    `json:"swift_bic"`
	Currency                     string    `json:"currency"`
	UnitName                     string    `json:"unit_name"`
	VatRate                      int       `json:"vat_rate"`
	DisplayedNote                string    `json:"displayed_note"`
	InvoiceNote                  string    `json:"invoice_note"`
	Due                          int       `json:"due"`
	InvoiceLanguage              string    `json:"invoice_language"`
	InvoicePaymentMethod         string    `json:"invoice_payment_method"`
	InvoiceProforma              bool      `json:"invoice_proforma"`
	InvoiceNumberFormat          string    `json:"invoice_number_format"`
	ProformaNumberFormat         string    `json:"proforma_number_format"`
	CustomEmailText              string    `json:"custom_email_text"`
	OverdueEmailText             string    `json:"overdue_email_text"`
	SendOverdueEmail             bool      `json:"send_overdue_email"`
	SendInvoiceFromProformaEmail bool      `json:"send_invoice_from_proforma_email"`
	SendThankYouEmail            bool      `json:"send_thank_you_email"`
	InvoicePaypal                bool      `json:"invoice_paypal"`
	InvoiceGopay                 bool      `json:"invoice_gopay"`
	Eet                          bool      `json:"eet"`
	EetInvoiceDefault            bool      `json:"eet_invoice_default"`
	HtmlUrl                      string    `json:"html_url"`
	Url                          string    `json:"url"`
	UpdatedAt                    time.Time `json:"updated_at"`
	CreatedAt                    time.Time `json:"created_at"`
}
