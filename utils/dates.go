package utils

import (
	"fmt"
	"time"
)

// GetMonthsStrings get the previous, current and next month as string in xx format. eg. 01 for January and 12 for December
func GetMonthStrings() (prevMonth, currentMonth, nextMonth string) {
	now := time.Now()

	// Current month as number
	current := now.Month()

	// Obtain year, previous and next month
	prev := now.AddDate(0, -1, 0).Month()
	next := now.AddDate(0, 1, 0).Month()

	// Format as string 01, 02, etc
	prevMonth = fmt.Sprintf("%02d", prev)
	currentMonth = fmt.Sprintf("%02d", current)
	nextMonth = fmt.Sprintf("%02d", next)

	return
}

func IsToday(dateToCompare time.Time) bool {
	today := time.Now().UTC()

	return dateToCompare.Year() == today.Year() &&
		dateToCompare.Month() == today.Month() &&
		dateToCompare.Day() == today.Day()
}

func ParseToUTC(dateStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05 -0700 MST"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %w", err)
	}
	return t.UTC(), nil
}
