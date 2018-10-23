package model

import (
	"encoding/json"
)

// ToJSON convert struct into string
func ToJSON(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}

// ListToJSON convert array of structs to JSON
func ListToJSON(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}
