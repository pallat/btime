package btime

import "time"

func AddMonth(t time.Time, months int) time.Time {
	a := t.Add(24 * time.Hour)
	b := a.AddDate(0, months, 0)
	c := b.Add(-24 * time.Hour)
	return c
}
