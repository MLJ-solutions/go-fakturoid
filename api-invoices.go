package go_fakturoid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MLJ-solutions/go-fakturoid/models"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const invoiceEndpoint = "invoices/%d.json"
const invoicesEndpoint = "invoices.json"

const (
	InvoiceQuerySince        = "since"
	InvoiceQueryUntil        = "until"
	InvoiceQueryUpdatedSince = "updated_since"
	InvoiceQueryUpdatedUntil = "updated_until"
	InvoiceQueryNumber       = "number"
	InvoiceQueryStatus       = "status"
	InvoiceQueryInvoiceId    = "invoice_id"
	InvoiceQueryCustomId     = "custom_id"
	InvoiceQuerySubjectId    = "subject_id"
)

// structure to help filter invoices
type InvoiceFilter struct {
	Since        *time.Time
	Until        *time.Time
	UpdatedSince *time.Time
	UpdatedUntil *time.Time
	Number       string
	Status       string
	InvoiceId    int
	SubjectId    int
	CustomId     string
}

func InvoiceFilterFromSubject(subject models.Subject) InvoiceFilter {
	return InvoiceFilter{SubjectId: subject.Id}
}

func (i InvoiceFilter) prepareMetadata() requestMetadata {
	metadata := requestMetadata{
		queryValues: url.Values{},
	}

	if i.Since != nil {
		metadata.queryValues.Add(InvoiceQuerySince, i.Since.String())
	}
	if i.Until != nil {
		metadata.queryValues.Add(InvoiceQueryUntil, i.Until.String())
	}
	if i.UpdatedSince != nil {
		metadata.queryValues.Add(InvoiceQueryUpdatedSince, i.UpdatedSince.String())
	}
	if i.UpdatedUntil != nil {
		metadata.queryValues.Add(InvoiceQueryUpdatedUntil, i.UpdatedUntil.String())
	}
	if i.Number != "" {
		metadata.queryValues.Add(InvoiceQueryNumber, i.Number)
	}
	if i.Status != "" {
		metadata.queryValues.Add(InvoiceQueryStatus, i.Status)
	}
	if i.InvoiceId != 0 {
		metadata.queryValues.Add(InvoiceQueryInvoiceId, strconv.Itoa(i.InvoiceId))
	}
	if i.CustomId != "" {
		metadata.queryValues.Add(InvoiceQueryCustomId, i.CustomId)
	}
	if i.SubjectId != 0 {
		metadata.queryValues.Add(InvoiceQuerySubjectId, strconv.Itoa(i.SubjectId))
	}

	return metadata
}

// request on /invoices.json
// TODO paging
func (c Client) Invoices(invoiceFilters InvoiceFilter) ([]models.Invoice, error) {
	metadata := invoiceFilters.prepareMetadata()

	resp, err := c.executeMethod(http.MethodGet, invoicesEndpoint, nil, metadata)
	if err != nil {
		return []models.Invoice{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return []models.Invoice{}, err
	}

	var Invoices []models.Invoice

	unmarshalErr := json.Unmarshal(body, &Invoices)
	//fmt.Print(string(body))
	if unmarshalErr != nil {
		return []models.Invoice{}, unmarshalErr
	}

	return Invoices, nil
}

const invoicesSearchEndpoint = "invoices/search.json"

// request on /Invoices/search.json?query=
func (c Client) InvoicesSearch(query string) ([]models.Invoice, error) {
	values := url.Values{}
	values.Add("query", query)

	resp, err := c.executeMethod(http.MethodGet, invoicesSearchEndpoint, nil, requestMetadata{
		queryValues: values,
	})
	if err != nil {
		return []models.Invoice{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return []models.Invoice{}, err
	}

	var Invoices []models.Invoice

	unmarshalErr := json.Unmarshal(body, &Invoices)
	//fmt.Print(string(body))
	if unmarshalErr != nil {
		return []models.Invoice{}, unmarshalErr
	}

	return Invoices, nil
}

// create Invoice on /invoices.json
func (c Client) CreateInvoice(Invoice models.Invoice) (models.Invoice, error) {
	requestBody, marshalErr := json.Marshal(Invoice)
	if marshalErr != nil {
		return models.Invoice{}, marshalErr
	}

	resp, err := c.executeMethodNoMeta(http.MethodPost, invoicesEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.Invoice{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.Invoice{}, err
	}

	unmarshalErr := json.Unmarshal(body, &Invoice)
	if unmarshalErr != nil {
		return models.Invoice{}, unmarshalErr
	}

	return Invoice, nil
}

// update Invoice on /invoices/{ID}.json
func (c Client) UpdateInvoice(Invoice models.Invoice, Id int) (models.Invoice, error) {
	requestBody, marshalErr := json.Marshal(Invoice)
	if marshalErr != nil {
		return models.Invoice{}, marshalErr
	}

	resp, err := c.executeMethodNoMeta(http.MethodPatch, fmt.Sprintf(invoiceEndpoint, Id), bytes.NewBuffer(requestBody))
	if err != nil {
		return models.Invoice{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.Invoice{}, err
	}

	unmarshalErr := json.Unmarshal(body, &Invoice)
	if unmarshalErr != nil {
		return models.Invoice{}, unmarshalErr
	}

	return Invoice, nil
}
