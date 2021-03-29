package go_fakturoid

import (
	"encoding/json"
	"github.com/MLJ-solutions/go-fakturoid/models"
	"io"
	"net/http"
)

const accountEndpoint = "account.json"

// reqeuest on /account.json
func (c Client) Account() (models.Account, error) {
	resp, err := c.executeMethodNoMeta(http.MethodGet, accountEndpoint, nil)
	if err != nil {
		return models.Account{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.Account{}, err
	}

	var accountInfo models.Account

	unmarshalErr := json.Unmarshal(body, &accountInfo)
	if unmarshalErr != nil {
		return models.Account{}, unmarshalErr
	}

	return accountInfo, nil
}
