package model

import (
	"time"
	"turns-app-go/utils"
)

type Reservation struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	OfficeID  string    `json:"office_id"`
}

type ReservationList struct {
	Reservations []Reservation `json:"turns"`
	Date         string        `json:"date"` // Date format "DD-MM-YYYY"
}

type WeekReservations struct {
	Monday    ReservationList `json:"monday"`
	Tuesday   ReservationList `json:"tuesday"`
	Wednesday ReservationList `json:"wednesday"`
	Thursday  ReservationList `json:"thursday"`
	Friday    ReservationList `json:"friday"`
	Saturday  ReservationList `json:"saturday"`
	Sunday    ReservationList `json:"sunday"`
}

func GetWeekReservationFromList(r []Reservation) WeekReservations {
	week := WeekReservations{}
	for _, reservation := range r {
		switch reservation.StartTime.Weekday() {
		case time.Monday:
			week.Monday.Reservations = append(week.Monday.Reservations, reservation)
		case time.Tuesday:
			week.Tuesday.Reservations = append(week.Tuesday.Reservations, reservation)
		case time.Wednesday:
			week.Wednesday.Reservations = append(week.Wednesday.Reservations, reservation)
		case time.Thursday:
			week.Thursday.Reservations = append(week.Thursday.Reservations, reservation)
		case time.Friday:
			week.Friday.Reservations = append(week.Friday.Reservations, reservation)
		case time.Saturday:
			week.Saturday.Reservations = append(week.Saturday.Reservations, reservation)
		case time.Sunday:
			week.Sunday.Reservations = append(week.Sunday.Reservations, reservation)
		}
	}

	// Add date to each day
	day := time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC)
	week_days := utils.GetWeekDates(day)
	week.Monday.Date = week_days[0]
	week.Tuesday.Date = week_days[1]
	week.Wednesday.Date = week_days[2]
	week.Thursday.Date = week_days[3]
	week.Friday.Date = week_days[4]
	week.Saturday.Date = week_days[5]
	week.Sunday.Date = week_days[6]

	return week
}
