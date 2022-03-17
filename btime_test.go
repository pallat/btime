package btime

import (
	"fmt"
	"testing"
	"time"
)

func TestAddMonths(t *testing.T) {
	date := time.Date(2013, time.January, 31, 23, 59, 59, 0, time.UTC)

	months := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	expected := [...]int{28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 31}

	for i, month := range months {
		d := AddMonth(date, month)
		fmt.Println(d)

		if d.Day() != expected[i] {
			t.Errorf("%d %v %v", i, d.Day(), expected[i])
		}
	}

}

func TestAddMonthsLeapYear(t *testing.T) {
	date := time.Date(2016, time.January, 31, 23, 59, 59, 0, time.UTC)

	months := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	expected := [...]int{29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 31}

	for i, month := range months {
		d := AddMonth(date, month)
		fmt.Println(d)

		if d.Day() != expected[i] {
			t.Errorf("%d %v %v", i, d.Day(), expected[i])
		}
	}

}
