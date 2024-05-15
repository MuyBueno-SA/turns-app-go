package utils

import (
	"testing"
	"time"
)

func assertTimeRange(t *testing.T, expected TimeRange, actual TimeRange) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetWeekByDay(t *testing.T) {

	weekRange := TimeRange{
		StartTime: time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 3, 3, 23, 59, 59, 0, time.UTC),
	}

	t.Run("test week mid day", func(t *testing.T) {
		day := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		week := GetWeekByDay(day)

		assertTimeRange(t, weekRange, week)
	})

	t.Run("test week monday", func(t *testing.T) {
		day := time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC)
		week := GetWeekByDay(day)

		assertTimeRange(t, weekRange, week)
	})

	t.Run("test week sunday", func(t *testing.T) {
		day := time.Date(2024, 3, 3, 23, 59, 59, 0, time.UTC)
		week := GetWeekByDay(day)

		assertTimeRange(t, weekRange, week)
	})
}

func TestGetWeekDates(t *testing.T) {
	t.Run("test week mid day", func(t *testing.T) {
		day := time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC)
		dates := GetWeekDates(day)

		expected := []string{"26-02-2024", "27-02-2024", "28-02-2024", "29-02-2024", "01-03-2024", "02-03-2024", "03-03-2024"}
		if len(dates) != len(expected) {
			t.Errorf("Expected %d dates, got %d", len(expected), len(dates))
		}
		for i, date := range dates {
			if date != expected[i] {
				t.Errorf("Expected %s, got %s", expected[i], date)
			}
		}
	})

	t.Run("test week monday", func(t *testing.T) {
		day := time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC)
		dates := GetWeekDates(day)

		expected := []string{"26-02-2024", "27-02-2024", "28-02-2024", "29-02-2024", "01-03-2024", "02-03-2024", "03-03-2024"}
		if len(dates) != len(expected) {
			t.Errorf("Expected %d dates, got %d", len(expected), len(dates))
		}
		for i, date := range dates {
			if date != expected[i] {
				t.Errorf("Expected %s, got %s", expected[i], date)
			}
		}
	})

	t.Run("test week sunday", func(t *testing.T) {
		day := time.Date(2024, 3, 3, 23, 59, 59, 0, time.UTC)
		dates := GetWeekDates(day)

		expected := []string{"26-02-2024", "27-02-2024", "28-02-2024", "29-02-2024", "01-03-2024", "02-03-2024", "03-03-2024"}
		if len(dates) != len(expected) {
			t.Errorf("Expected %d dates, got %d", len(expected), len(dates))
		}
		for i, date := range dates {
			if date != expected[i] {
				t.Errorf("Expected %s, got %s", expected[i], date)
			}
		}
	})
}
