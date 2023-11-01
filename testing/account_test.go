package testing

import (
	"net/http"
	"testing"
)

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	if http.StatusOK != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}
}
