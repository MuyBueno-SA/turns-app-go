package main

import (
	"net/http"
	"time"
	db "turns-app-go/dbmanager"
	"turns-app-go/helper"
	"turns-app-go/server"
	"turns-app-go/utils"
)

const ConfigPath = "configs/app_config.dev.toml"

func main() {

	config, _ := utils.LoadConfig(ConfigPath)
	today := time.Now()
	usersDBManager := db.DefaultInMemoryUsersDBManager()
	reservationsDBManager := &db.InMemoryReservationsDBManager{
		Reservations: helper.GenerateRandomTurns(20, config.Business, today),
	}

	dbManager := db.DBManager{UsersManager: usersDBManager, ReservationsManager: reservationsDBManager}
	s := server.NewAPPServer(dbManager, config.Business)
	http.ListenAndServe(":5000", s)
}
