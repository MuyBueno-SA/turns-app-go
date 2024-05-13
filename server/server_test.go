package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeartbeat(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
	response := httptest.NewRecorder()

	Heartbeat(response, request)

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "I'm alive!")
}

func TestNewAPPServer(t *testing.T) {
	s := NewAPPServer(nil)

	request, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
	response := httptest.NewRecorder()

	s.ServeHTTP(response, request)

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "I'm alive!")
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
