package btime

import "time"

var monthDays = [...]int{
	0,
	31,
	28,
	31,
	30,
	31,
	30,
	31,
	31,
	30,
	31,
	30,
	31,
}

func AddMonth(t time.Time, months int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	if day > monthDays[month+time.Month(months)%12] {
		day = monthDays[month+time.Month(months)]
	}

	return time.Date(year, month+time.Month(months), day, hour, min, sec, 0, t.Location())
}
