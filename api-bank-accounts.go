package go_fakturoid

import (
	"encoding/json"
	"github.com/MLJ-solutions/go-fakturoid/models"
	"io/ioutil"
	"net/http"
)

const bankAccountsEndpoint = "bank_accounts.json"

// request on /bank_accounts.json
func (c Client) BankAccounts() ([]models.BankAccount, error) {
	resp, err := c.executeMethodNoMeta(http.MethodGet, bankAccountsEndpoint, nil)
	if err != nil {
		return []models.BankAccount{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return []models.BankAccount{}, err
	}

	var bankAccounts []models.BankAccount

	unmarshalErr := json.Unmarshal(body, &bankAccounts)
	//fmt.Print(string(body))
	if unmarshalErr != nil {
		return []models.BankAccount{}, unmarshalErr
	}

	return bankAccounts, nil
}
