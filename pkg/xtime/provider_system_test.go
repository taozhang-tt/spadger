package xtime

import (
	"testing"
	"time"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestSystemProvider(t *testing.T) {
	Convey("Sleep", t, func() {
		beg := Now()
		Sleep(100 * time.Millisecond)
		v := Now().Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 100*time.Millisecond)
		So(v, ShouldAlmostEqual, 100*time.Millisecond, 10*time.Millisecond)
	})

	Convey("After", t, func() {
		beg := Now()
		<-After(100 * time.Millisecond)
		v := Now().Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 100*time.Millisecond)
		So(v, ShouldAlmostEqual, 100*time.Millisecond, 10*time.Millisecond)
	})

	Convey("AfterFunc", t, func() {
		beg := Now()
		ch := make(chan time.Time, 1)
		AfterFunc(100*time.Millisecond, func() { ch <- Now() })
		end := <-ch
		v := end.Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 100*time.Millisecond)
		So(v, ShouldAlmostEqual, 100*time.Millisecond, 10*time.Millisecond)
	})

	Convey("NewTimer", t, func() {
		beg := Now()
		timer := NewTimer(100 * time.Millisecond)
		end := <-timer.Chan()
		v := end.Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 100*time.Millisecond)
		So(v, ShouldAlmostEqual, 100*time.Millisecond, 10*time.Millisecond)
	})

	Convey("TimerStop", t, func() {
		timer := NewTimer(50 * time.Millisecond)
		active := timer.Stop()
		So(active, ShouldBeTrue)
		time.Sleep(time.Millisecond)
		select {
		case <-timer.Chan():
			So("should not be here", ShouldBeNil)
		default:
		}
	})

	Convey("TimerReset", t, func() {
		beg := Now()
		timer := NewTimer(300 * time.Millisecond)
		timer.Reset(100 * time.Millisecond)
		end := <-timer.Chan()
		v := end.Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 100*time.Millisecond)
		So(v, ShouldAlmostEqual, 100*time.Millisecond, 10*time.Millisecond)
	})

	Convey("NewTicker", t, func() {
		beg := Now()
		ticker := NewTicker(100 * time.Millisecond)

		end := <-ticker.Chan()
		v := end.Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 100*time.Millisecond)
		So(v, ShouldAlmostEqual, 100*time.Millisecond, 10*time.Millisecond)

		end = <-ticker.Chan()
		v = end.Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 200*time.Millisecond)
		So(v, ShouldAlmostEqual, 200*time.Millisecond, 20*time.Millisecond)

		ticker.Stop()
		time.Sleep(time.Millisecond)
		select {
		case <-ticker.Chan():
			So("should not be here", ShouldBeNil)
		default:
		}
	})

	Convey("Tick", t, func() {
		beg := Now()
		ch := Tick(100 * time.Millisecond)

		end := <-ch
		v := end.Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 100*time.Millisecond)
		So(v, ShouldAlmostEqual, 100*time.Millisecond, 10*time.Millisecond)

		end = <-ch
		v = end.Sub(beg)
		So(v, ShouldBeGreaterThanOrEqualTo, 200*time.Millisecond)
		So(v, ShouldAlmostEqual, 200*time.Millisecond, 20*time.Millisecond)

		So(Tick(0), ShouldBeNil)
	})
}
