package dbmanager

import (
	"errors"
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

func (db *InMemoryReservationsDBManager) AddReservation(turn model.Reservation) error {
	if db.isTurnAvailable(turn) {
		db.Reservations = append(db.Reservations, turn)
		return nil
	}
	return errors.New("turn not available")
}

func (db *InMemoryReservationsDBManager) isTurnAvailable(turn model.Reservation) bool {
	for _, reservation := range db.Reservations {
		// if office is different continue the loop
		if reservation.OfficeID != turn.OfficeID {
			continue
		}
		// if day is different continue the loop
		if reservation.StartTime.Year() != turn.StartTime.Year() || reservation.StartTime.YearDay() != turn.StartTime.YearDay() {
			continue
		}

		// Starts before and ends after
		if turn.StartTime.Before(reservation.StartTime) && turn.EndTime.After(reservation.EndTime) {
			return false
		}
		// Starts inside range
		if turn.StartTime.After(reservation.StartTime) && turn.StartTime.Before(reservation.EndTime) {
			return false
		}
		// Ends inside range
		if turn.EndTime.After(reservation.StartTime) && turn.EndTime.Before(reservation.EndTime) {
			return false
		}
	}
	return true
}

func DefaultInMemoryUsersDBManager() *InMemoryUsersDBManager {
	return &InMemoryUsersDBManager{Users: model.UsersSlice}
}

func DefaultInMemoryReservationsDBManager() *InMemoryReservationsDBManager {
	return &InMemoryReservationsDBManager{Reservations: model.ReservationsSlice}
}
