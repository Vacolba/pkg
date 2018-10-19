package errors

import "testing"

func TestValidateError(t *testing.T) {
	err := NewValidateError()
	err.AddError("p1", "error p1")
	err.AddError("p2", "error p2")
	if err.Errors["p1"] != "error p1" {
		t.Fatal(err.Errors["p1"])
	}
	if err.Errors["p2"] != "error p2" {
		t.Fatal(err.Errors["p1"])
	}

	if err.Error() != "Errors: p1, p2" {
		t.Fatal("Los errores no coinciden")
	}

	if err.ToJSON() != "{\"Errors\":{\"p1\":\"error p1\",\"p2\":\"error p2\"}}" {
		t.Fatal(err.ToJSON())
	}
}
