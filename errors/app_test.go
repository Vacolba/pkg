package errors

import "testing"

func TestNewLocAppError(t *testing.T) {
	err := NewLocAppError("TestAppError", "test_app_error", nil, "error explain")

	if err.ID != "test_app_error" {
		t.Fatal("ID error: ", err.ID)
	}
	if err.Message != "test_app_error" {
		t.Fatal("Message error: ", err.Message)
	}

	if err.Where != "TestAppError" {
		t.Fatal("Where error: ", err.Where)
	}

	if err.DetailedError != "error explain" {
		t.Fatal("DetailedError error: ", err.DetailedError)
	}

	if err.Error() != "TestAppError: test_app_error, error explain" {
		t.Fatal(err.Error())
	}

	if err.ToJSON() != "{\"id\":\"test_app_error\",\"message\":\"test_app_error\",\"detailed_error\":\"error explain\",\"Params\":null}" {
		t.Fatal(err.ToJSON())
	}
}

func TestNewAppError(t *testing.T) {
	err := NewAppError("TestAppError", "test_app_error", nil, "error explain", 200)
	if err.Error() != "TestAppError: test_app_error, error explain" {
		t.Fatal(err.Error())
	}

	if err.ToJSON() != "{\"id\":\"test_app_error\",\"message\":\"test_app_error\",\"detailed_error\":\"error explain\",\"status_code\":200,\"Params\":null}" {
		t.Fatal(err.ToJSON())
	}
}
