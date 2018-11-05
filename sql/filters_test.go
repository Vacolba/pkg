package sql

import "testing"

func TestFilter(t *testing.T) {
	f := NewFilter()

	s, c := f.ToSQL()
	if s != "WHERE 1 = 0" {
		t.Fatal(f.ToSQL())
	}
	f.Add("id", "=", "1")
	if f.String() != "WHERE id = 1" {
		t.Fatal(f.String())
	}

	s, c = f.ToSQL()
	if s != "WHERE id = ?" || c[0] != "1" {
		t.Fatal(f.ToSQL())
	}

	f.Add("active", "=", "0")
	if f.String() != "WHERE id = 1\nAND active = 0" {
		t.Fatal(f.String())
	}

	s, c = f.ToSQL()
	if s != "WHERE id = ?\nAND active = ?" || c[0] != "1" || c[1] != "0" {
		t.Fatal(f.ToSQL())
	}
}

func TestFilterError(t *testing.T) {
	f := NewFilter()
	err := f.Add("id", "ts", "1")
	if err == nil {
		t.Fatal("Should fail")
	}
}

func TestFilterDateRanges(t *testing.T) {
	f := NewFilter()
	f.AddDateRange("created_at", "2017-01-01", "2017-01-31")
	if f.String() != "WHERE created_at BETWEEN 2017-01-01 00:00:00 AND 2017-01-31 23:59:59" {
		t.Fatal(f.String())
	}

	s, c := f.ToSQL()
	if s != "WHERE created_at BETWEEN ? AND ?" || c[0] != "2017-01-01 00:00:00" || c[1] != "2017-01-31 23:59:59" {
		t.Fatal(f.ToSQL())
	}

	f.Add("id", "=", "1")
	if f.String() != "WHERE created_at BETWEEN 2017-01-01 00:00:00 AND 2017-01-31 23:59:59\nAND id = 1" {
		t.Fatal(f.String())
	}

}
