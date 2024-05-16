package helper

import (
	"math/rand"
	"time"

	dbm "turns-app-go/dbmanager"
	"turns-app-go/model"
	"turns-app-go/utils"
)

func hourFromTimeString(timeString string) int {
	hour, _ := time.Parse("15.04", timeString)
	return hour.Hour()
}

// generate random turns in the day's week with no overlaping
func GenerateRandomTurns(max int, business utils.BusinessConfig, day time.Time, reservationsDB *dbm.InMemoryReservationsDBManager) {
	weekDays := utils.GetWeekDates(day)
	modules := []int{}
	for i := hourFromTimeString(business.StartTime); i < hourFromTimeString(business.EndTime); i++ {
		modules = append(modules, i)
	}

	for i := 0; i < max; i++ {
		randomUserPosition := rand.Intn(len(model.UsersSlice))
		randomUser := model.UsersSlice[randomUserPosition]
		randomOfficePosition := rand.Intn(len(business.Offices))
		randomOffice := business.Offices[randomOfficePosition]

		randomDayPosition := rand.Intn(len(weekDays))
		randomDayString := weekDays[randomDayPosition]
		randomDay, _ := time.Parse("02-01-2006", randomDayString)

		randomStartTimePosition := rand.Intn(len(modules))
		randomStartTime := modules[randomStartTimePosition]

		randomDuration := rand.Intn(4) + 1
		randomEndTime := randomStartTime + randomDuration
		if randomEndTime > hourFromTimeString(business.EndTime) {
			randomEndTime = hourFromTimeString(business.EndTime)
		}

		turn := model.Reservation{
			ID:        i,
			UserID:    randomUser.ID,
			StartTime: time.Date(randomDay.Year(), randomDay.Month(), randomDay.Day(), randomStartTime, 0, 0, 0, time.UTC),
			EndTime:   time.Date(randomDay.Year(), randomDay.Month(), randomDay.Day(), randomEndTime, 0, 0, 0, time.UTC),
			OfficeID:  randomOffice,
		}
		err := reservationsDB.AddReservation(turn)
		if err != nil {
			i--
			continue
		}

	}
}
