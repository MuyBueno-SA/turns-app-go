package utils

import (
	"testing"
	"time"
)

func CompareTime(t *testing.T, expected time.Time, actual time.Time) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetWeekByDay(t *testing.T) {
	week_start := time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC)
	week_end := time.Date(2024, 3, 3, 23, 59, 59, 0, time.UTC)

	day := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)

	week := GetWeekByDay(day)

	CompareTime(t, week_start, week.StartTime)
	CompareTime(t, week_end, week.EndTime)

}
