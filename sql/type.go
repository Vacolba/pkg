package sql

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

// NullString easy work with sql.NullString
type NullString struct {
	sql.NullString
}

// NewNullString init a new NullString
func NewNullString(data string) NullString {
	ns := new(NullString)
	ns.Scan(data)
	return *ns
}

// UnmarshalJSON returns NullString object from a JSON
func (ns *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		ns.String = ""
		ns.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &ns.String)
	if err != nil {
		return err
	}
	ns.Valid = true
	return nil
}

// MarshalJSON returns JSON representation of NullString
func (ns NullString) MarshalJSON() (b []byte, err error) {
	if ns.String == "" && !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// NullInt64 easy work with sql.NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// UnmarshalJSON returns NullInt64 object from a JSON
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		ni.Int64 = 0
		ni.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &ni.Int64)
	if err != nil {
		return err
	}
	ni.Valid = true
	return nil
}

func (ni NullInt64) MarshalJSON() (b []byte, err error) {
	if ni.Int64 == 0 && !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// NullType is an Struct to work with sql Null DateTime.
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
