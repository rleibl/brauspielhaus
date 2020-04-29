package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type MyTime struct {
	Day   int
	Month int
	Year  int
}

func ParseTime(s string) MyTime {

	m := MyTime{
		Day:   1,
		Month: 1,
		Year:  1970,
	}

	parts := strings.Split(s, ".")

	v, err := strconv.Atoi(parts[0])
	if err != nil {
		// FIXME log error
		return m
	}
	m.Day = v

	v, err = strconv.Atoi(parts[1])
	if err != nil {
		// FIXME log error
		return m
	}
	m.Month = v

	v, err = strconv.Atoi(parts[2])
	if err != nil {
		// FIXME log error
		return m
	}
	m.Year = v

	return m
}

func (m *MyTime) ToString() string {

	r := fmt.Sprintf("%02d.%02d.%d", m.Day, m.Month, m.Year)

	return r
}

func (m *MyTime) DurationBetween(m2 MyTime) int {

	t1 := time.Date(m.Year, IntToMonth(m.Month), m.Day, 1, 0, 0, 0, time.UTC)
	t2 := time.Date(m2.Year, IntToMonth(m2.Month), m2.Day, 1, 0, 0, 0, time.UTC)

	var d time.Duration
	if t1.After(t2) {
		d = t1.Sub(t2)
	} else {
		d = t2.Sub(t1)
	}

	days := int(d.Hours() / 24)

	return days
}

func IntToMonth(i int) time.Month {

	switch i {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	case 12:
		return time.December
	default:
		return time.January // FIXME fail here
	}
}
