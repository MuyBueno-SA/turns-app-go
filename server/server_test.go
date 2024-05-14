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

func TestGetUsers(t *testing.T) {
	dbManager := db.DBManager{UsersManager: getInMemoryUsersDBManager()}
	s := APPServer{DBManager: dbManager}

	t.Run("GET /users", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		response := httptest.NewRecorder()

		s.GetUsers(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		want := getUser0()
		got := getUsersResponse(t, response.Body)
		if !reflect.DeepEqual(got[0], want) {
			t.Errorf("got %v, want %v", got[0], want)
		}
	})
}

func TestNewAPPServer(t *testing.T) {

	dbManager := db.DBManager{UsersManager: getInMemoryUsersDBManager()}
	s := NewAPPServer(dbManager)

	t.Run("GET /users", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		response := httptest.NewRecorder()

		s.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		want := getUser0()
		got := getUsersResponse(t, response.Body)
		if !reflect.DeepEqual(got[0], want) {
			t.Errorf("got %v, want %v", got[0], want)
		}
	})

	t.Run("GET /heartbeat", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
		response := httptest.NewRecorder()

		s.ServeHTTP(response, request)

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

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func getInMemoryUsersDBManager() *db.InMemoryUsersDBManager {
	users_slice := []model.User{
		getUser0(),
		getUser2(),
		getUser4(),
	}
	return &db.InMemoryUsersDBManager{
		Users: users_slice,
	}
}

func getUser0() model.User {
	return model.User{
		ID:       0,
		Username: "Virginia D'Espósito",
		Email:    "vir@test.com",
		Phone:    "123456789",
		Activity: "Psicopedagoga",
	}
}

func getUser2() model.User {
	return model.User{
		ID:       2,
		Username: "Federico Bogado",
		Email:    "fico@test.com",
		Phone:    "123456789",
		Activity: "Developer",
	}
}

func getUser4() model.User {
	return model.User{
		ID:       4,
		Username: "Susana Horia",
		Email:    "susana@other.com",
		Phone:    "987654321",
		Activity: "Reiki",
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
