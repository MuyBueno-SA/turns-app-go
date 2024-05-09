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
	weekday := int(day.Weekday())

	mYear, mMonth, mDay := day.AddDate(0, 0, -weekday+1).Date()
	start := time.Date(mYear, mMonth, mDay, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 6).Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	return TimeRange{
		StartTime: start,
		EndTime:   end,
	}
}
