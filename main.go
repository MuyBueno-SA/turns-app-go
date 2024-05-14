package main

import (
	"net/http"
	"os"
	db "turns-app-go/dbmanager"
	"turns-app-go/server"
)

func main() {
	file, _ := os.Getwd()
	dbUsersFilePath := file + "/testing_files/users.json"
	usersDBManager := db.NewInMemoryUsersDBManager(dbUsersFilePath)
	dbManager := db.DBManager{UsersManager: usersDBManager}
	s := server.NewAPPServer(dbManager)
	http.ListenAndServe(":5000", s)
}
