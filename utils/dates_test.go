package utils

import (
	"testing"
	"time"
)

func TestGetMonths(t *testing.T) {
	previous, current, next := GetMonthStrings()

	if previous != "03" || current != "04" || next != "05" {
		t.Fatal("invalid dates")
	}
}

func TestComparesIfDateIsToday(t *testing.T) {
	date1, err := ParseToUTC("2025-04-20 19:00:00 +0000 UTC")
	if err != nil {
		t.Fatal(err)
	}

	date2, err := ParseToUTC("2025-05-03 19:00:00 +0000 UTC")
	if err != nil {
		t.Fatal(err)
	}

	date3, err := ParseToUTC("2025-04-16 19:00:00 +0000 UTC")
	if err != nil {
		t.Fatal(err)
	}

	if IsToday(date1) || IsToday(date2) || IsToday(date3) {
		t.Fatal("error, expecting to be false")
	}

	if !IsToday(time.Now()) {
		t.Fatal("expected to be true")
	}

}
