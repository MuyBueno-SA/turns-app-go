package dbmanager

import (
	"turns-app-go/model"
)

type InMemoryUsersDBManager struct {
	Users []model.User
}

func (db *InMemoryUsersDBManager) GetUsers() []model.User {
	return db.Users
}

type InMemoryReservationsDBManager struct {
	Reservations []model.Reservation
}

func (db *InMemoryReservationsDBManager) GetWeek(date string) model.WeekReservations {
	week := model.GetWeekReservationFromList(db.Reservations)
	return week
}

func DefaultInMemoryUsersDBManager() *InMemoryUsersDBManager {
	return &InMemoryUsersDBManager{Users: model.UsersSlice}
}

func DefaultInMemoryReservationsDBManager() *InMemoryReservationsDBManager {
	return &InMemoryReservationsDBManager{Reservations: model.ReservationsSlice}
}
