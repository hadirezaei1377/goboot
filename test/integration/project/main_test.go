package project

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsersHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsersHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	expected := `[{"id":1,"name":"Ali"},{"id":2,"name":"Bardia"}]`
	assert.JSONEq(t, expected, rr.Body.String(), "Response body should match")
}
