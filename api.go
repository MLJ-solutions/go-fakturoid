package go_fakturoid

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// List of success status.
var successStatus = []int{
	http.StatusOK,
	http.StatusNoContent,
	http.StatusPartialContent,
}

func (c Client) executeMethod(method string, endpoint string, body io.Reader, metadata requestMetadata) (res *http.Response, err error) {
	// create target url with selected accountEndpoint
	targetUrl, err := c.constructUrl(endpoint, metadata.queryValues)
	if err != nil {
		return nil, err
	}

	// create request
	req, err := http.NewRequest(method, targetUrl.String(), body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		errRequest := ToErrorResponse(err)
		return nil, errRequest
	}

	//fmt.Println(req)
	res, err = c.do(req) // todo handle
	if err != nil {
		return nil, err
	}
	//fmt.Println("res")
	//fmt.Println(res)

	// For any known successful http status, return quickly.
	for _, httpStatus := range successStatus {
		if httpStatus == res.StatusCode {
			return res, nil
		}
	}

	//log.Panic(res)
	all, err := ioutil.ReadAll(res.Body)
	closeResponse(res)
	if err != nil {
		return nil, err
	}

	var apiError ErrorResponse
	json.Unmarshal(all, &apiError)
	log.Println(string(all))
	log.Println(apiError.Errors)

	return nil, apiError

}

func (c Client) constructUrl(endpoint string, queryValues url.Values) (*url.URL, error) {
	urlStr := c.EndpointURL().String() + endpoint

	// If there are any query values, add them to the end.
	if len(queryValues) > 0 {
		urlStr = urlStr + "?" + queryValues.Encode()
	}
	return url.Parse(urlStr)
}
