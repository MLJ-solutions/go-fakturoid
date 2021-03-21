package go_fakturoid

import (
	"io"
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
	if err != nil {
		errResponse := ToErrorResponse(err)
		return nil, errResponse
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

	//// TODO Read the body to be saved later.
	//errBodyBytes, err := ioutil.ReadAll(res.Body)
	//// res.Body should be closed
	//closeResponse(res)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil

}

func (c Client) constructUrl(endpoint string, queryValues url.Values) (*url.URL, error) {
	urlStr := c.EndpointURL().String() + endpoint

	// If there are any query values, add them to the end.
	if len(queryValues) > 0 {
		urlStr = urlStr + "?" + queryValues.Encode()
	}
	return url.Parse(urlStr)
}
