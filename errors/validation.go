package errors

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ValidateError is an Struct that contains validation errors
type ValidateError struct {
	Errors map[string]interface{}
}

// NewValidateError creates an returns an ValidateError Struct
func NewValidateError() *ValidateError {
	v := ValidateError{}
	v.Errors = make(map[string]interface{})
	return &v
}

// AddError add a new error
func (er *ValidateError) AddError(id string, value interface{}) {
	er.Errors[id] = value
}

// Error returns a string representing the error
func (er *ValidateError) Error() string {
	keys := reflect.ValueOf(er.Errors).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return "Errors: " + strings.Join(strkeys, ", ")
}

// ToJSON returns the error representative string in JSON format
func (er *ValidateError) ToJSON() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	}
	return string(b)
}
