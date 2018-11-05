package sql

import (
	"errors"
	"strings"
)

const (
	QueryTypes = map[string]bool{
		"=",
		"!=",
		"in",
		"not in",
		"like",
		"no like",
	}
)

// Condition is each one of the conditions of the consultation
type Condition struct {
	Field string
	Type  string
	Value interface{}
}

func (c *Condition) String() string {
	return c.Field + " " + c.Type + " " + c.Value.(string)
}

func (c *Condition) ToSQL() (string, interface{}) {
	return c.Field + " " + c.Type + " ?", c.Value
}

type DateRangeCondition struct {
	Field string
	From  string
	To    string
}

func (c *DateRangeCondition) String() string {
	fFrom := c.From + " 00:00:00"
	fTo := c.To + " 23:59:59"
	return c.Field + " BETWEEN " + fFrom + " AND " + fTo
}

func (c *DateRangeCondition) ToSQL() (string, interface{}, interface{}) {
	fFrom := c.From + " 00:00:00"
	fTo := c.To + " 23:59:59"
	return c.Field + " BETWEEN ? AND ?", fFrom, fTo
}

// Filter is an Struct that contains the query filters
type Filter struct {
	Conditions          []*Condition
	DateRangeConditions []*DateRangeCondition
}

// NewFilter returns a new Filter Struct
func NewFilter() *Filter {
	f := Filter{}
	return &f
}

func (f *Filter) Add(field, qType string, value interface{}) error {
	if !QueryTypes[qType] {
		return errors.New("No valid Query type")
	}
	c := &Condition{Field: field, Type: qType, Value: value}
	f.Conditions = append(f.Conditions, c)
	return nil
}

func (f *Filter) AddDateRange(field, from, to string) {
	c := &DateRangeCondition{Field: field, From: from, To: to}
	f.DateRangeConditions = append(f.DateRangeConditions, c)
}

func (f *Filter) String() string {
	l := 0
	t := make([]string, len(f.Conditions)+len(f.DateRangeConditions))
	for i := range f.DateRangeConditions {
		t[l] = f.DateRangeConditions[i].String()
		l++
	}
	for i := range f.Conditions {
		t[l] = f.Conditions[i].String()
		l++
	}

	return "WHERE " + strings.Join(t, "\nAND ")
}

func (f *Filter) ToSQL() (string, []interface{}) {
	c := make([]interface{}, len(f.Conditions)+(len(f.DateRangeConditions)*2))
	if (len(f.Conditions) + len(f.DateRangeConditions)) < 1 {
		return "WHERE 1 = 0", c
	}

	l := 0
	r := 0
	t := make([]string, len(f.Conditions)+len(f.DateRangeConditions))
	for i := range f.DateRangeConditions {
		t[l], c[r], c[r+1] = f.DateRangeConditions[i].ToSQL()
		l++
		r += 2
	}

	for i := range f.Conditions {
		t[l], c[r] = f.Conditions[i].ToSQL()
		l++
		r++
	}
	return "WHERE " + strings.Join(t, "\nAND "), c
}
