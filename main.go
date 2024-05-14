package main

import (
	"net/http"
	db "turns-app-go/dbmanager"
	"turns-app-go/server"
)

func main() {
	usersDBManager := db.DefaultInMemoryUsersDBManager()
	dbManager := db.DBManager{UsersManager: usersDBManager}
	s := server.NewAPPServer(dbManager)
	http.ListenAndServe(":5000", s)
}
