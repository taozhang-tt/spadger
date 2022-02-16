package xtime

import (
	"testing"
	"time"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestTime(t *testing.T) {
	// 2020-08-20 02:13:14 Thu
	n := time.Date(2020, 8, 20, 2, 13, 14, 0, time.Local)

	Convey("Time", t, func() {
		So(FormatTime(With(n).BeginningOfHour()), ShouldEqual, "2020-08-20 02:00:00")
		So(FormatTime(With(n).EndOfHour()), ShouldEqual, "2020-08-20 02:59:59")
		So(FormatTime(With(n).BeginningOfDay()), ShouldEqual, "2020-08-20 00:00:00")
		So(FormatTime(With(n).EndOfDay()), ShouldEqual, "2020-08-20 23:59:59")
		So(FormatTime(With(n).BeginningOfWeek()), ShouldEqual, "2020-08-16 00:00:00") // Sun
		So(FormatTime(With(n).EndOfWeek()), ShouldEqual, "2020-08-22 23:59:59")       // Sat
		So(FormatTime(With(n).BeginningOfMonth()), ShouldEqual, "2020-08-01 00:00:00")
		So(FormatTime(With(n).EndOfMonth()), ShouldEqual, "2020-08-31 23:59:59")
		So(FormatTime(With(n).BeginningOfQuarter()), ShouldEqual, "2020-07-01 00:00:00")
		So(FormatTime(With(n).EndOfQuarter()), ShouldEqual, "2020-09-30 23:59:59")
		So(FormatTime(With(n).BeginningOfYear()), ShouldEqual, "2020-01-01 00:00:00")
		So(FormatTime(With(n).EndOfYear()), ShouldEqual, "2020-12-31 23:59:59")
	})

	Convey("Time.Config", t, func() {
		c := &Config{
			DayStartHour: 5,
			WeekStartDay: time.Monday,
		}
		So(FormatTime(c.With(n).BeginningOfHour()), ShouldEqual, "2020-08-20 02:00:00")
		So(FormatTime(c.With(n).EndOfHour()), ShouldEqual, "2020-08-20 02:59:59")
		So(FormatTime(c.With(n).BeginningOfDay()), ShouldEqual, "2020-08-19 05:00:00")
		So(FormatTime(c.With(n).EndOfDay()), ShouldEqual, "2020-08-20 04:59:59")
		So(FormatTime(c.With(n).BeginningOfWeek()), ShouldEqual, "2020-08-17 05:00:00")
		So(FormatTime(c.With(n).EndOfWeek()), ShouldEqual, "2020-08-24 04:59:59")
		So(FormatTime(c.With(n).BeginningOfMonth()), ShouldEqual, "2020-08-01 05:00:00")
		So(FormatTime(c.With(n).EndOfMonth()), ShouldEqual, "2020-09-01 04:59:59")
		So(FormatTime(c.With(n).BeginningOfQuarter()), ShouldEqual, "2020-07-01 05:00:00")
		So(FormatTime(c.With(n).EndOfQuarter()), ShouldEqual, "2020-10-01 04:59:59")
		So(FormatTime(c.With(n).BeginningOfYear()), ShouldEqual, "2020-01-01 05:00:00")
		So(FormatTime(c.With(n).EndOfYear()), ShouldEqual, "2021-01-01 04:59:59")

		So(FormatTime(c.With(MustParseTime("2020-08-16 03:01:01")).BeginningOfWeek()), ShouldEqual, "2020-08-10 05:00:00")
		So(FormatTime(c.With(MustParseTime("2020-08-16 10:01:01")).BeginningOfWeek()), ShouldEqual, "2020-08-10 05:00:00")
		So(FormatTime(c.With(MustParseTime("2020-08-17 03:01:01")).BeginningOfWeek()), ShouldEqual, "2020-08-10 05:00:00")
		So(FormatTime(c.With(MustParseTime("2020-08-17 10:01:01")).BeginningOfWeek()), ShouldEqual, "2020-08-17 05:00:00")
	})
}
