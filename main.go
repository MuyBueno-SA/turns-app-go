package main

import (
	"net/http"
	db "turns-app-go/dbmanager"
	"turns-app-go/server"
	"turns-app-go/utils"
)

const ConfigPath = "configs/app_config.dev.toml"

func main() {
	usersDBManager := db.DefaultInMemoryUsersDBManager()
	config, _ := utils.LoadConfig(ConfigPath)

	dbManager := db.DBManager{UsersManager: usersDBManager}
	s := server.NewAPPServer(dbManager, config.Business)
	http.ListenAndServe(":5000", s)
}
