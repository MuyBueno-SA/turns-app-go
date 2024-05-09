package utils

import (
	"testing"
	"time"
)

func TestGetWeekByDay(t *testing.T) {
	week_start := time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC)
	week_end := time.Date(2024, 3, 3, 23, 59, 59, 0, time.UTC)

	day := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)

	week := GetWeekByDay(day)

	if week.StartTime != week_start {
		t.Errorf("Expected %v, got %v", week_start, week.StartTime)
	}

	if week.EndTime != week_end {
		t.Errorf("Expected %v, got %v", week_end, week.EndTime)
	}
}
