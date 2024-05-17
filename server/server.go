package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	dbm "turns-app-go/dbmanager"
	"turns-app-go/model"
	"turns-app-go/utils"
)

const jsonContentType = "application/json"

type BusinessInfo struct {
	BusinessConfig utils.BusinessConfig `json:"business_config"`
	Users          []model.User         `json:"users"`
}

// JSONResponse writes the data as a JSON response to the given http.ResponseWriter.
// It sets the Content-Type header to "application/json".
func JSONResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(data)
}

// TODO: This can't be ok
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// APPServer is the main server for the application.
type APPServer struct {
	BusinessConfig utils.BusinessConfig
	DBManager      dbm.DBManager
	http.Handler
}

// GetBusinessInfo returns the business information.
// It returns the business configuration and the list of users.
// The response is in JSON format.
func (s *APPServer) GetBusinessInfo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	businessInfo := BusinessInfo{
		BusinessConfig: s.BusinessConfig,
		Users:          s.DBManager.UsersManager.GetUsers(),
	}
	JSONResponse(w, businessInfo)
}

// GetUsers returns the list of users.
// The response is in JSON format.
func (s *APPServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	users := s.DBManager.UsersManager.GetUsers()
	JSONResponse(w, users)
}

// GetWeek returns the reservations for the week of the given date.
// The response is in JSON format.
// The date is expected to be in the format "dd.mm.yyyy".
func (s *APPServer) GetWeek(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	date := r.URL.Query().Get("date")
	week := s.DBManager.ReservationsManager.GetWeek(date)
	JSONResponse(w, week)
}

// Heartbeat is a simple handler that returns "I'm alive!".
func Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alive!")
}

// NewAPPServer creates a new APPServer with the given DBManager and BusinessConfig.
// It sets up the routes for the server.
func NewAPPServer(manager dbm.DBManager, config utils.BusinessConfig) *APPServer {
	s := new(APPServer)

	s.DBManager = manager
	s.BusinessConfig = config
	router := http.NewServeMux()

	router.Handle("/heartbeat", http.HandlerFunc(Heartbeat))
	router.Handle("/users", http.HandlerFunc(s.GetUsers))
	router.Handle("/business_info", http.HandlerFunc(s.GetBusinessInfo))
	router.Handle("/get_week", http.HandlerFunc(s.GetWeek))
	s.Handler = router

	return s
}
