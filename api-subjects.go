package go_fakturoid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MLJ-solutions/go-fakturoid/models"
	"io"
	"net/http"
	"net/url"
)

const subjectsEndpoint = "subjects.json"

// request on /subjects.json
// TODO paging
func (c Client) Subjects() ([]models.Subject, error) {
	resp, err := c.executeMethod(http.MethodGet, subjectsEndpoint, nil, requestMetadata{})
	if err != nil {
		return []models.Subject{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return []models.Subject{}, err
	}

	var Subjects []models.Subject

	unmarshalErr := json.Unmarshal(body, &Subjects)
	//fmt.Print(string(body))
	if unmarshalErr != nil {
		return []models.Subject{}, unmarshalErr
	}

	return Subjects, nil
}

const subjectsSearchEndpoint = "subjects/search.json"

// request on /subjects/search.json?query=
func (c Client) SubjectsSearch(query string) ([]models.Subject, error) {
	values := url.Values{}
	values.Add("query", query)

	resp, err := c.executeMethod(http.MethodGet, subjectsSearchEndpoint, nil, requestMetadata{
		queryValues: values,
	})
	if err != nil {
		return []models.Subject{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return []models.Subject{}, err
	}

	var Subjects []models.Subject

	unmarshalErr := json.Unmarshal(body, &Subjects)
	//fmt.Print(string(body))
	if unmarshalErr != nil {
		return []models.Subject{}, unmarshalErr
	}

	return Subjects, nil
}

// create subject on /subjects.json
func (c Client) CreateSubject(subject models.Subject) (models.Subject, error) {
	requestBody, marshalErr := json.Marshal(subject)
	if marshalErr != nil {
		return models.Subject{}, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, subjectsEndpoint, bytes.NewBuffer(requestBody), requestMetadata{})
	if err != nil {
		return models.Subject{}, err
	}

	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.Subject{}, err
	}

	unmarshalErr := json.Unmarshal(body, &subject)
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return models.Subject{}, unmarshalErr
	}

	return subject, nil
}
