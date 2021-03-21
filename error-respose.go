package go_fakturoid

import (
	"encoding/xml"
	"fmt"
)

// TODO refine
type ErrorResponse struct {
	XMLName xml.Name `xml:"Error" json:"-"`
	Code    string
	Message string
}

func ToErrorResponse(err error) ErrorResponse {
	switch err := err.(type) {
	case ErrorResponse:
		return err
	default:
		return ErrorResponse{}
	}
}

// Error - Returns S3 error string.
func (e ErrorResponse) Error() string {
	if e.Message == "" {
		msg, ok := errorResponseMap[e.Code]
		if !ok {
			msg = fmt.Sprintf("Error response code %s.", e.Code)
		}
		return msg
	}
	return e.Message
}
