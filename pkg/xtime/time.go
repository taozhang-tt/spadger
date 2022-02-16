package xtime

import "time"

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
	Day                       = 24 * Hour
	Week                      = 7 * Day
)

var uptime = Now().Truncate(time.Second)

func Uptime() time.Time {
	return uptime
}

var DefaultConfig = &Config{
	DayStartHour: 0,
	WeekStartDay: time.Sunday,
}

type Config struct {
	DayStartHour int
	WeekStartDay time.Weekday
}

func (c *Config) With(t time.Time) *Time {
	return &Time{Time: t, Config: c}
}

type Time struct {
	time.Time
	*Config
}

func With(t time.Time) *Time {
	return DefaultConfig.With(t)
}

func (t Time) BeginningOfHour() time.Time {
	h := t.Time.Hour()
	y, m, d := t.Time.Date()
	return time.Date(y, m, d, h, 0, 0, 0, t.Location())
}

func (t Time) BeginningOfDay() time.Time {
	h := t.Time.Hour()
	y, m, d := t.Time.Date()
	if h >= t.DayStartHour {
		return time.Date(y, m, d, t.DayStartHour, 0, 0, 0, t.Location())
	} else {
		return time.Date(y, m, d-1, t.DayStartHour, 0, 0, 0, t.Location())
	}
}

func (t Time) BeginningOfWeek() time.Time {
	v := t.BeginningOfDay()
	w := int(v.Weekday())
	if t.WeekStartDay != time.Sunday {
		ws := int(t.WeekStartDay)
		if w < ws {
			w = w - ws + 7
		} else {
			w = w - ws
		}
	}
	return v.AddDate(0, 0, -w)
}

func (t Time) BeginningOfMonth() time.Time {
	y, m, _ := t.Time.Date()
	return time.Date(y, m, 1, t.DayStartHour, 0, 0, 0, t.Location())
}

func (t Time) BeginningOfQuarter() time.Time {
	m := t.BeginningOfMonth()
	offset := (int(m.Month()) - 1) % 3
	return m.AddDate(0, -offset, 0)
}

func (t Time) BeginningOfYear() time.Time {
	return time.Date(t.Time.Year(), time.January, 1, t.DayStartHour, 0, 0, 0, t.Location())
}

func (t Time) EndOfHour() time.Time {
	return t.BeginningOfHour().Add(Hour - time.Nanosecond)
}

func (t Time) EndOfDay() time.Time {
	return t.BeginningOfDay().Add(Day - time.Nanosecond)
}

func (t Time) EndOfWeek() time.Time {
	return t.BeginningOfWeek().Add(Week - time.Nanosecond)
}

func (t Time) EndOfMonth() time.Time {
	return t.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (t Time) EndOfQuarter() time.Time {
	return t.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

func (t Time) EndOfYear() time.Time {
	return t.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}
