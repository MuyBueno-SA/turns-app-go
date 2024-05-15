package utils

import (
	"time"
)

type TimeRange struct {
	StartTime time.Time
	EndTime   time.Time
}

// GetWeekByDay returns the start and end of the week for the given day.
// Start is Monday 00:00:00 and end is Sunday 23:59:59.
func GetWeekByDay(day time.Time) TimeRange {
	time_zone := day.Location()
	weekday := int(day.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	mYear, mMonth, mDay := day.AddDate(0, 0, -weekday+1).Date()
	start := time.Date(mYear, mMonth, mDay, 0, 0, 0, 0, time_zone)
	end := start.AddDate(0, 0, 6).Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	return TimeRange{
		StartTime: start,
		EndTime:   end,
	}
}

func GetWeekDates(day time.Time) []string {
	weekday := int(day.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	mYear, mMonth, mDay := day.AddDate(0, 0, -weekday+1).Date()
	start := time.Date(mYear, mMonth, mDay, 0, 0, 0, 0, time.UTC)

	dates := []string{}
	for i := 0; i < 7; i++ {
		date := start.AddDate(0, 0, i)
		dates = append(dates, date.Format("02-01-2006"))
	}

	return dates

}
