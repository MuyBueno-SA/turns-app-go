package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	dbm "turns-app-go/dbmanager"
)

type APPServer struct {
	DBManager dbm.DBManager
	http.Handler
}

func (s *APPServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.DBManager.UsersManager.GetUsers()
	json.NewEncoder(w).Encode(users)
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alive!")
}

func NewAPPServer(manager dbm.DBManager) *APPServer {
	s := new(APPServer)

	s.DBManager = manager
	router := http.NewServeMux()

	router.Handle("/heartbeat", http.HandlerFunc(Heartbeat))
	router.Handle("/users", http.HandlerFunc(s.GetUsers))
	s.Handler = router

	return s
}
