package go_fakturoid

import (
	"bytes"
	"fmt"
	"strings"
)

// TODO refine
type ErrorResponse struct {
	Errors map[string][]string
}

func ToErrorResponse(err error) ErrorResponse {
	switch err := err.(type) {
	case ErrorResponse:
		return err
	default:
		return ErrorResponse{}
	}
}

// Error - Returns error string.
func (e ErrorResponse) Error() string {
	b := new(bytes.Buffer)
	for key, value := range e.Errors {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, strings.Join(value, ", "))
	}

	return b.String()
}
