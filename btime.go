package btime

import "time"

func AddMonth(t time.Time, months int) time.Time {
	return t.Add(24*time.Hour).AddDate(0, months, 0).Add(-24 * time.Hour)
}
