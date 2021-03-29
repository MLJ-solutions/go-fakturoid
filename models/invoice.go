package models

import "log"

type InvoiceLine struct {
	Name      string  `json:"name"`
	Quantity  float32 `json:"quantity,string"`
	UnitName  string  `json:"unit_name"`
	UnitPrice float32 `json:"unit_price,string"`
	VatRate   int     `json:"vat_rate"`
}

type Invoice struct {
	Id                      int               `json:"id,omitempty"`
	CustomId                string            `json:"custom_id,omitempty"`
	Proforma                bool              `json:"proforma,omitempty"`
	PartialProforma         bool              `json:"partial_proforma,omitempty"`
	Number                  string            `json:"number,omitempty"`
	VariableSymbol          string            `json:"variable_symbol,omitempty"`
	YourName                string            `json:"your_name,omitempty"`
	YourStreet              string            `json:"your_street,omitempty"`
	YourStreet2             string            `json:"your_street2,omitempty"`
	YourCity                string            `json:"your_city,omitempty"`
	YourZip                 string            `json:"your_zip,omitempty"`
	YourCountry             string            `json:"your_country,omitempty"`
	YourRegistrationNo      string            `json:"your_registration_no,omitempty"`
	YourVatNo               string            `json:"your_vat_no,omitempty"`
	YourLocalVatNo          string            `json:"your_local_vat_no,omitempty"`
	ClientName              string            `json:"client_name,omitempty"`
	ClientStreet            string            `json:"client_street,omitempty"`
	ClientStreet2           string            `json:"client_street2,omitempty"`
	ClientCity              string            `json:"client_city,omitempty"`
	ClientZip               string            `json:"client_zip,omitempty"`
	ClientCountry           string            `json:"client_country,omitempty"`
	ClientRegistrationNo    string            `json:"client_registration_no,omitempty"`
	ClientVatNo             string            `json:"client_vat_no,omitempty"`
	ClientLocalVatNo        string            `json:"client_local_vat_no,omitempty"`
	SubjectId               int               `json:"subject_id,omitempty"`
	SubjectCustomId         int               `json:"subject_custom_id,omitempty"`
	GeneratorId             int               `json:"generator_id,omitempty"`
	RelatedId               int               `json:"related_id,omitempty"`
	Correction              bool              `json:"correction,omitempty"`
	CorrectionId            int               `json:"correction_id,omitempty"`
	Paypal                  bool              `json:"paypal,omitempty"`
	Gopay                   bool              `json:"gopay,omitempty"`
	Token                   string            `json:"token,omitempty"`
	Status                  string            `json:"status,omitempty"`
	OrderNumber             string            `json:"order_number,omitempty"`
	IssuedOn                FakturoidDate     `json:"issued_on,omitempty"`
	TaxableFulfillmentDue   FakturoidDate     `json:"taxable_fulfillment_due,omitempty"`
	Due                     int               `json:"due,omitempty"`
	DueOn                   FakturoidDate     `json:"due_on,omitempty"`
	SentAt                  FakturoidDateTime `json:"sent_at,omitempty"`
	PaidAt                  FakturoidDateTime `json:"paid_at,omitempty"`
	ReminderSentAt          FakturoidDate     `json:"reminder_sent_at,omitempty"`
	AcceptedAt              FakturoidDateTime `json:"accepted_at,omitempty"`
	CancelledAt             FakturoidDateTime `json:"cancelled_at,omitempty"`
	WebInvoiceSeenAt        FakturoidDateTime `json:"webinvoice_seen_at,omitempty"`
	Note                    string            `json:"note,omitempty"`
	FooterNote              string            `json:"footer_note,omitempty"`
	PrivateNote             string            `json:"private_note,omitempty"`
	Tags                    []string          `json:"tags,omitempty"`
	BankAccountId           int               `json:"bank_account_id,omitempty"`
	BankAccount             string            `json:"bank_account,omitempty"`
	Iban                    string            `json:"iban,omitempty"`
	SwiftBic                string            `json:"swift_bic,omitempty"`
	PaymentMethod           string            `json:"payment_method,omitempty"`
	HideBankAccount         bool              `json:"hide_bank_account,omitempty"`
	Currency                string            `json:"currency,omitempty"`
	ExchangeRate            string            `json:"exchange_rate,omitempty"`
	Language                string            `json:"language,omitempty"`
	TransferredTaxLiability bool              `json:"transferred_tax_liability,omitempty"`
	EuElectronicService     bool              `json:"eu_electronic_service,omitempty"`
	VatPriceMode            string            `json:"vat_price_mode,omitempty"`
	SupplyCode              string            `json:"supply_code,omitempty"`
	Subtotal                float32           `json:"subtotal,string,omitempty"`
	Total                   float32           `json:"total,string,omitempty"`
	NativeSubtotal          float32           `json:"native_subtotal,string,omitempty"`
	NativeTotal             float32           `json:"native_total,string,omitempty"`
	RemainingAmount         float32           `json:"remaining_amount,string,omitempty"`
	RemainingNativeAmount   float32           `json:"remaining_native_amount,string,omitempty"`
	PaidAmount              float32           `json:"paid_amount,string,omitempty"`
	Lines                   []InvoiceLine     `json:"lines"`
	Attachment              string            `json:"attachment,omitempty"`
	HtmlUrl                 string            `json:"html_url,omitempty"`
	PublicHtmlUrl           string            `json:"public_html_url,omitempty"`
	Url                     string            `json:"url,omitempty"`
	PdfUrl                  string            `json:"pdf_url,omitempty"`
	SubjectUrl              string            `json:"subject_url,omitempty"`
	CreatedAt               string            `json:"created_at,omitempty"`
	UpdatedAt               string            `json:"updated_at,omitempty"`

	// TODO EET
	//Eet                     bool        `json:"eet,omitempty"`
	//EetCashRegister         string        `json:"eet_cash_register,omitempty"`
	//EetStore                string        `json:"eet_store,omitempty"`
	//EetRecords              string        `json:"eet_records,omitempty"`
}

const (
	EventMarkAsSent         = "mark_as_sent"
	EventDeliver            = "deliver"
	EventPay                = "pay"
	EventPayProforma        = "pay_proforma"
	EventPayPartialProforma = "pay_partial_proforma"
	EventRemovePayment      = "remove_payment"
	EventDeliverReminder    = "deliver_reminder"
	EventCancel             = "cancel"
	EventUndoCancel         = "undo_cancel"
	EventLock               = "lock"
	EventUnlock             = "unlock"
)

type InvoiceEvent struct {
	Event          string
	PaidAt         *FakturoidDateTime
	PaidAmount     float32
	VariableSymbol string
	BankAccountId  int
}

func NewPayInvoiceEvent(PaidAt *FakturoidDateTime, PaidAmount float32, VariableSymbol string, BankAccountId int) *InvoiceEvent {
	return &InvoiceEvent{
		Event:          EventPay,
		PaidAt:         PaidAt,
		PaidAmount:     PaidAmount,
		VariableSymbol: VariableSymbol,
		BankAccountId:  BankAccountId,
	}
}

func NewInvoiceEvent(event string) *InvoiceEvent {
	switch event {
	case EventMarkAsSent:
		return &InvoiceEvent{Event: EventMarkAsSent}
	case EventDeliver:
		return &InvoiceEvent{Event: EventDeliver}
	case EventPay:
		return &InvoiceEvent{Event: EventPay}
	case EventPayProforma:
		return &InvoiceEvent{Event: EventPayProforma}
	case EventPayPartialProforma:
		return &InvoiceEvent{Event: EventPayPartialProforma}
	case EventRemovePayment:
		return &InvoiceEvent{Event: EventRemovePayment}
	case EventDeliverReminder:
		return &InvoiceEvent{Event: EventDeliverReminder}
	case EventCancel:
		return &InvoiceEvent{Event: EventCancel}
	case EventUndoCancel:
		return &InvoiceEvent{Event: EventUndoCancel}
	case EventLock:
		return &InvoiceEvent{Event: EventLock}
	case EventUnlock:
		return &InvoiceEvent{Event: EventUnlock}
	}

	log.Println("unrecognised event")
	return &InvoiceEvent{Event: event}
}
