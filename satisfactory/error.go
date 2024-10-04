package satisfactory

import (
	"fmt"
)

// ErrorResponse represents an error returned by the API.
type ErrorResponse struct {
	Response     *Response
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	ErrorData    interface{} `json:"errorData"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("API Error %s: %s", e.ErrorCode, e.ErrorMessage)
}
