// TODO This data doesn't belong here

package model

import (
	"time"
)

var UsersSlice = []User{
	{
		ID:       0,
		Username: "Virginia D'Espósito",
		Email:    "vir@test.com",
		Phone:    "123456789",
		Activity: "Psicopedagoga",
	},
	{
		ID:       2,
		Username: "Federico Bogado",
		Email:    "fico@test.com",
		Phone:    "123456789",
		Activity: "Developer",
	},
	{
		ID:       4,
		Username: "Susana Horia",
		Email:    "susana@other.com",
		Phone:    "987654321",
		Activity: "Reiki",
	},
}

var ReservationsSlice = []Reservation{
	{
		ID:        0,
		UserID:    0,
		StartTime: time.Date(2024, 2, 26, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 2, 26, 10, 0, 0, 0, time.UTC),
		OfficeID:  "OFF_01",
	},
	{
		ID:        1,
		UserID:    0,
		StartTime: time.Date(2024, 2, 26, 10, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 2, 26, 11, 0, 0, 0, time.UTC),
		OfficeID:  "OFF_01",
	},
	{
		ID:        2,
		UserID:    0,
		StartTime: time.Date(2024, 2, 26, 11, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 2, 26, 13, 0, 0, 0, time.UTC),
		OfficeID:  "OFF_02",
	},
	{
		ID:        3,
		UserID:    2,
		StartTime: time.Date(2024, 2, 26, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 2, 26, 10, 0, 0, 0, time.UTC),
		OfficeID:  "OFF_02",
	},
	{
		ID:        4,
		UserID:    2,
		StartTime: time.Date(2024, 2, 27, 10, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 2, 27, 11, 0, 0, 0, time.UTC),
		OfficeID:  "OFF_01",
	},
}

func GetDefaultWeekReservations() WeekReservations {
	week := WeekReservations{}
	week.Monday = ReservationList{Reservations: []Reservation{
		ReservationsSlice[0],
		ReservationsSlice[1],
		ReservationsSlice[2],
		ReservationsSlice[3],
	}, Date: "26-02-2024"}
	week.Tuesday = ReservationList{Reservations: []Reservation{
		ReservationsSlice[4],
	}, Date: "27-02-2024"}
	week.Wednesday = ReservationList{Reservations: []Reservation{}, Date: "28-02-2024"}
	week.Thursday = ReservationList{Reservations: []Reservation{}, Date: "29-02-2024"}
	week.Friday = ReservationList{Reservations: []Reservation{}, Date: "01-03-2024"}
	week.Saturday = ReservationList{Reservations: []Reservation{}, Date: "02-03-2024"}
	week.Sunday = ReservationList{Reservations: []Reservation{}, Date: "03-03-2024"}
	return week
}
