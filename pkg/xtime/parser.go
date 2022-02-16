package xtime

import "time"

const (
	DateLayout        = "2006-01-02"
	TimeLayout        = "2006-01-02 15:04:05"
	RFC3339Layout     = "2006-01-02T15:04:05"
	RFC3339NanoLayout = "2006-01-02T15:04:05.000"
)

func ParseTime(s string) (time.Time, error) {
	return time.ParseInLocation(TimeLayout, s, time.Local)
}

func MustParseTime(s string) time.Time {
	t, err := ParseTime(s)
	if err != nil {
		panic(err)
	}
	return t
}

func FormatTime(t time.Time) string {
	return t.Format(TimeLayout)
}

func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation(DateLayout, s, time.Local)
}

func MustParseDate(s string) time.Time {
	t, err := ParseDate(s)
	if err != nil {
		panic(err)
	}
	return t
}

func FormatDate(t time.Time) string {
	return t.Format(DateLayout)
}

func RFC3339() string {
	return Now().Format(RFC3339Layout)
}

func RFC3339Nano() string {
	return Now().Format(RFC3339NanoLayout)
}
