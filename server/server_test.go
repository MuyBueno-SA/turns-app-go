package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	db "turns-app-go/dbmanager"
	"turns-app-go/model"
)

func TestHeartbeat(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
	response := httptest.NewRecorder()

	Heartbeat(response, request)

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "I'm alive!")
}

func TestAPPServer(t *testing.T) {
	dbManager := db.DBManager{UsersManager: db.DefaultInMemoryUsersDBManager()}
	s := APPServer{DBManager: dbManager}

	t.Run("test GetUsers", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		response := httptest.NewRecorder()

		s.GetUsers(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)

		users := getUsersResponse(t, response.Body)
		assertUsers(t, users, db.UsersSlice)
	})
}

func TestNewAPPServerRouting(t *testing.T) {

	dbManager := db.DBManager{UsersManager: db.DefaultInMemoryUsersDBManager()}
	server := NewAPPServer(dbManager)

	t.Run("GET /users", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		users := getUsersResponse(t, response.Body)
		assertUsers(t, users, db.UsersSlice)
	})

	t.Run("GET /heartbeat", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "I'm alive!")
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertUsers(t *testing.T, got, want []model.User) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func getUsersResponse(t *testing.T, body io.Reader) (users []model.User) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&users)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of User, '%v'", body, err)
	}
	return
}
