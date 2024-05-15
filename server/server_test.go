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
	"turns-app-go/utils"
)

const ConfigPath = "../configs/app_test_config.toml"

func TestHeartbeat(t *testing.T) {
	request, response := prepareHeartbeatRequest()
	Heartbeat(response, request)
	assertHeartbeatRequest(t, response)
}

func TestAPPServer(t *testing.T) {
	dbManager := db.DBManager{UsersManager: db.DefaultInMemoryUsersDBManager(),
		ReservationsManager: db.DefaultInMemoryReservationsDBManager()}
	config, _ := utils.LoadConfig(ConfigPath)
	s := APPServer{DBManager: dbManager, BusinessConfig: config.Business}

	t.Run("test BusinessInfo", func(t *testing.T) {
		request, response := prepareBusinessInfoRequest()
		s.GetBusinessInfo(response, request)
		assertBusinessInfoRequest(t, response)
	})

	t.Run("test GetUsers", func(t *testing.T) {
		request, response := prepareGetUsersRequest()
		s.GetUsers(response, request)
		assertGetUsersRequest(t, response)
	})

	t.Run("test GetWeek", func(t *testing.T) {
		request, response := prepareGetWeekRequest()
		s.GetWeek(response, request)
		assertGetWeekRequest(t, response)
	})
}

func TestNewAPPServerRouting(t *testing.T) {

	dbManager := db.DBManager{UsersManager: db.DefaultInMemoryUsersDBManager(),
		ReservationsManager: db.DefaultInMemoryReservationsDBManager()}
	config, _ := utils.LoadConfig(ConfigPath)
	server := NewAPPServer(dbManager, config.Business)

	t.Run("GET /heartbeat", func(t *testing.T) {
		request, response := prepareHeartbeatRequest()
		server.ServeHTTP(response, request)
		assertHeartbeatRequest(t, response)
	})

	t.Run("GET /users", func(t *testing.T) {
		request, response := prepareGetUsersRequest()
		server.ServeHTTP(response, request)
		assertGetUsersRequest(t, response)
	})

	t.Run("GET /business_info", func(t *testing.T) {
		request, response := prepareBusinessInfoRequest()
		server.ServeHTTP(response, request)
		assertBusinessInfoRequest(t, response)
	})

	t.Run("GET /get_week", func(t *testing.T) {
		request, response := prepareGetWeekRequest()
		server.ServeHTTP(response, request)
		assertGetWeekRequest(t, response)
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

func assertHeartbeatRequest(t *testing.T, response *httptest.ResponseRecorder) {
	t.Helper()
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "I'm alive!")
}

func assertGetUsersRequest(t *testing.T, response *httptest.ResponseRecorder) {
	t.Helper()
	assertStatus(t, response.Code, http.StatusOK)
	assertContentType(t, response, jsonContentType)
	users := getUsersResponse(t, response.Body)
	assertUsers(t, users, model.UsersSlice)
}

func assertBusinessInfoRequest(t *testing.T, response *httptest.ResponseRecorder) {
	t.Helper()
	assertStatus(t, response.Code, http.StatusOK)
	assertContentType(t, response, jsonContentType)
	business := getBusinessInfoResponse(t, response.Body)
	assertBusinessInfo(t, business, getTestBusinessInfo())
}

func assertUsers(t *testing.T, got, want []model.User) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertBusinessInfo(t *testing.T, got, want BusinessInfo) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertWeek(t *testing.T, got, want model.WeekReservations) {
	t.Helper()
	if got.Monday.Date != want.Monday.Date {
		t.Errorf("got %v want %v", got.Monday.Date, want.Monday.Date)
	}
	if got.Tuesday.Date != want.Tuesday.Date {
		t.Errorf("got %v want %v", got.Tuesday.Date, want.Tuesday.Date)
	}
	if got.Sunday.Date != want.Sunday.Date {
		t.Errorf("got %v want %v", got.Sunday.Date, want.Sunday.Date)
	}

	if len(got.Monday.Reservations) != len(want.Monday.Reservations) {
		t.Errorf("got %v want %v", len(got.Monday.Reservations), len(want.Monday.Reservations))
	}
	if len(got.Tuesday.Reservations) != len(want.Tuesday.Reservations) {
		t.Errorf("got %v want %v", len(got.Tuesday.Reservations), len(want.Tuesday.Reservations))
	}
	if len(got.Sunday.Reservations) != len(want.Sunday.Reservations) {
		t.Errorf("got %v want %v", len(got.Sunday.Reservations), len(want.Sunday.Reservations))
	}

}

func assertGetWeekRequest(t *testing.T, response *httptest.ResponseRecorder) {
	t.Helper()
	assertStatus(t, response.Code, http.StatusOK)
	assertContentType(t, response, jsonContentType)
	week := getWeekResponse(t, response.Body)
	assertWeek(t, week, model.GetDefaultWeekReservations())
}

func getUsersResponse(t *testing.T, body io.Reader) (users []model.User) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&users)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of User, '%v'", body, err)
	}
	return
}

func getBusinessInfoResponse(t *testing.T, body io.Reader) (business BusinessInfo) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&business)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into BusinessConfig, '%v'", body, err)
	}
	return
}

func getWeekResponse(t *testing.T, body io.Reader) (week model.WeekReservations) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&week)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into WeekReservations, '%v'", body, err)
	}
	return
}

func prepareHeartbeatRequest() (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(http.MethodGet, "/heartbeat", nil)
	response := httptest.NewRecorder()
	return request, response
}

func prepareGetUsersRequest() (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	response := httptest.NewRecorder()
	return request, response
}

func prepareBusinessInfoRequest() (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(http.MethodGet, "/business_info", nil)
	response := httptest.NewRecorder()
	return request, response
}

func prepareGetWeekRequest() (*http.Request, *httptest.ResponseRecorder) {

	request, _ := http.NewRequest(http.MethodGet, "/get_week?date=2024-02-28", nil)
	response := httptest.NewRecorder()
	return request, response
}

func getTestBusinessConfig() utils.BusinessConfig {
	return utils.BusinessConfig{
		Name:          "Test Business",
		StartTime:     "08.00",
		EndTime:       "21.00",
		MinModuleTime: 60,
		Offices:       []string{"OFF_01", "OFF_02"},
	}
}

func getTestBusinessInfo() BusinessInfo {
	return BusinessInfo{
		BusinessConfig: getTestBusinessConfig(),
		Users:          model.UsersSlice,
	}
}
