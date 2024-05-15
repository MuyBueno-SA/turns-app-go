package model

import (
	"testing"
)

func TestGetWeekReservationFromList(t *testing.T) {
	reservations := ReservationsSlice
	week := GetWeekReservationFromList(reservations)
	if len(week.Monday.Reservations) != 4 {
		t.Errorf("Expected 4 reservations on Monday, got %d", len(week.Monday.Reservations))
	}
	if len(week.Tuesday.Reservations) != 1 {
		t.Errorf("Expected 1 reservation on Tuesday, got %d", len(week.Tuesday.Reservations))
	}
	if week.Monday.Reservations[0].ID != 0 {
		t.Errorf("Expected reservation ID 0 on Monday, got %d", week.Monday.Reservations[0].ID)
	}
}
