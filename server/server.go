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

func JSONResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(data)
}

type APPServer struct {
	BusinessConfig utils.BusinessConfig
	DBManager      dbm.DBManager
	http.Handler
}

func (s *APPServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.DBManager.UsersManager.GetUsers()
	JSONResponse(w, users)
}

func (s *APPServer) GetBusinessInfo(w http.ResponseWriter, r *http.Request) {
	businessInfo := BusinessInfo{
		BusinessConfig: s.BusinessConfig,
		Users:          s.DBManager.UsersManager.GetUsers(),
	}
	JSONResponse(w, businessInfo)
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alive!")
}

func NewAPPServer(manager dbm.DBManager, config utils.BusinessConfig) *APPServer {
	s := new(APPServer)

	s.DBManager = manager
	s.BusinessConfig = config
	router := http.NewServeMux()

	router.Handle("/heartbeat", http.HandlerFunc(Heartbeat))
	router.Handle("/users", http.HandlerFunc(s.GetUsers))
	router.Handle("/business_info", http.HandlerFunc(s.GetBusinessInfo))
	s.Handler = router

	return s
}
