package sql

import (
	"testing"
	"time"
)

func TestNullString(t *testing.T) {
	ns := NullString{}
	if ns.String != "" || ns.Valid {
		t.Fatal("Not init empty")
	}

	b, _ := ns.MarshalJSON()
	if string(b) != "null" {
		t.Fatal("MarshalJSON not null")
	}

	ns.Scan("Test Done")

	b, _ = ns.MarshalJSON()
	if !ns.Valid && string(b) != "\"Test Done\"" {
		t.Fatal("MarshalJSON not same")
	}

	ns.UnmarshalJSON([]byte("null"))
	if ns.String != "" || ns.Valid {
		t.Fatal("Not Marshaled to empty")
	}

	err := ns.UnmarshalJSON([]byte("\"Test Done\""))
	if !ns.Valid || string(b) != "\"Test Done\"" || err != nil {
		t.Log(err)
		t.Fatal("Not Marshaled to string")
	}

	ns = NullString{}
	err = ns.UnmarshalJSON([]byte(`true`))
	if ns.Valid {
		t.Fatal("Must fail")
	}
	nns := NewNullString("test")
	if nns.String != "test" {
		t.Fatal("No correct creation")
	}
}

func TestNullInt64(t *testing.T) {
	ni := NullInt64{}
	if ni.Int64 != 0 || ni.Valid {
		t.Fatal("Not init empty")
	}

	b, _ := ni.MarshalJSON()
	if string(b) != "null" {
		t.Fatal("MarshalJSON not null")
	}

	ni.Scan(58)

	b, _ = ni.MarshalJSON()
	if !ni.Valid && string(b) != "58" {
		t.Fatal("MarshalJSON not same")
	}

	ni.UnmarshalJSON([]byte("null"))
	if ni.Int64 != 0 || ni.Valid {
		t.Fatal("Not Marshaled to empty")
	}

	err := ni.UnmarshalJSON([]byte("58"))
	if !ni.Valid || string(b) != "58" || err != nil {
		t.Log(err)
		t.Fatal("Not Marshaled to string")
	}

	ni = NullInt64{}
	err = ni.UnmarshalJSON([]byte(":)"))
	if ni.Valid {
		t.Fatal("Must fail")
	}
}

func TestNullTime(t *testing.T) {
	nt := NullTime{}
	v, _ := nt.Value()
	if v != nil {
		t.Fatal("Not init empty")
	}
	tn := time.Now()
	nt.Scan(tn)
	if !nt.Valid {
		t.Fatal("Not scaned to Valid")
	}
	stn, _ := nt.Value()
	if stn != tn {
		t.Fatal("Not scaned time")
	}
}
