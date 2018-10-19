package errors

import (
	"encoding/json"
)

// AppError is an Struct that contains error info.
type AppError struct {
	ID            string `json:"id"`
	Message       string `json:"message"`               // Message to be display to the end user without debugging information
	DetailedError string `json:"detailed_error"`        // Internal error string to help the developer
	RequestID     string `json:"request_id,omitempty"`  // The RequestId that's also set in the header
	StatusCode    int    `json:"status_code,omitempty"` // The http status code, should be a const in https://golang.org/pkg/net/http/
	Where         string `json:"-"`                     // The function where it happened in the form of Struct.Func
	Params        map[string]interface{}
}

// NewAppError creates an returns an AppError Struct
func NewAppError(where string, id string, params map[string]interface{}, details string, status int) *AppError {
	ap := NewLocAppError(where, id, params, details)
	ap.StatusCode = status
	return ap
}

// NewLocAppError creates an returns an AppError Struct for internal error
func NewLocAppError(where string, id string, params map[string]interface{}, details string) *AppError {
	ap := &AppError{}
	ap.ID = id
	ap.Params = params
	ap.Message = id
	ap.Where = where
	ap.DetailedError = details
	return ap
}

// Error returns a string representing the error
func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

// ToJSON returns the error representative string in JSON format
func (er *AppError) ToJSON() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	}
	return string(b)
}
