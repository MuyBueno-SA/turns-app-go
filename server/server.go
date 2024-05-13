package server

import (
	"fmt"
	"net/http"
	dbm "turns-app-go/dbmanager"
)

type APPServer struct {
	dbManager dbm.DBManager
	http.Handler
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alive!")
}

func NewAPPServer(dbManager dbm.DBManager) *APPServer {
	p := new(APPServer)

	p.dbManager = dbManager
	router := http.NewServeMux()

	router.Handle("/heartbeat", http.HandlerFunc(Heartbeat))
	p.Handler = router

	return p
}
