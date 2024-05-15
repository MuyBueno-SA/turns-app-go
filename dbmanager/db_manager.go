package dbmanager

import (
	"turns-app-go/model"
)

type usersManagerI interface {
	GetUsers() []model.User
}

type reservationsManagerI interface {
	GetWeek(date string) model.WeekReservations
}

type DBManager struct {
	UsersManager        usersManagerI
	ReservationsManager reservationsManagerI
}
