package xtime

import (
	"container/heap"
	"math/rand"
	"runtime"
	"testing"
	"time"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestFrozenProvider(t *testing.T) {
	Freeze()
	defer Unfreeze()

	Convey("Sleep", t, func() {
		for i := 0; i < 50; i++ {
			ch := make(chan time.Duration)
			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			beg := Now()
			go func() {
				Sleep(delay)
				ch <- delay
			}()

			So(WaitSched(1, time.Second), ShouldBeTrue)

			Advance(50 * time.Millisecond)
			runtime.Gosched()
			select {
			case <-ch:
				So("should not be here", ShouldBeNil)
			default:
			}
			Advance(delay - 50*time.Millisecond)

			v := <-ch
			So(v, ShouldEqual, delay)
			So(Now().Sub(beg), ShouldEqual, delay)
		}
	})

	Convey("After", t, func() {
		for i := 0; i < 50; i++ {
			ch := make(chan time.Time)
			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			beg := Now()
			go func() {
				v := <-After(delay)
				ch <- v
			}()

			So(WaitSched(1, time.Second), ShouldBeTrue)

			Advance(50 * time.Millisecond)
			runtime.Gosched()
			select {
			case <-ch:
				So("should not be here", ShouldBeNil)
			default:
			}
			Advance(delay - 50*time.Millisecond)

			v := <-ch
			So(v, ShouldEqual, Now())
			So(Now().Sub(beg), ShouldEqual, delay)
		}
	})

	Convey("AfterFunc", t, func() {
		for i := 0; i < 50; i++ {
			ch := make(chan time.Duration)
			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			beg := Now()

			AfterFunc(delay, func() { ch <- delay })

			So(WaitSched(1, time.Second), ShouldBeTrue)

			Advance(50 * time.Millisecond)
			runtime.Gosched()
			select {
			case <-ch:
				So("should not be here", ShouldBeNil)
			default:
			}
			Advance(delay - 50*time.Millisecond)

			v := <-ch
			So(v, ShouldEqual, delay)
			So(Now().Sub(beg), ShouldEqual, delay)
		}
	})

	Convey("NewTimer", t, func() {
		for i := 0; i < 50; i++ {
			ch := make(chan time.Time)
			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			beg := Now()

			timer := NewTimer(delay)
			go func() {
				v := <-timer.Chan()
				ch <- v
			}()

			So(WaitSched(1, time.Second), ShouldBeTrue)

			Advance(50 * time.Millisecond)
			runtime.Gosched()
			select {
			case <-ch:
				So("should not be here", ShouldBeNil)
			default:
			}
			Advance(delay - 50*time.Millisecond)

			v := <-ch
			So(v, ShouldEqual, Now())
			So(Now().Sub(beg), ShouldEqual, delay)

			// the timer must has already expired or been stopped.
			So(timer.Stop(), ShouldBeFalse)
		}
	})

	Convey("TimerStop", t, func() {
		for i := 0; i < 50; i++ {
			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			timer := NewTimer(delay)
			So(WaitSched(1, time.Second), ShouldBeTrue)
			So(timer.Stop(), ShouldBeTrue)
			Advance(delay + time.Second)
			select {
			case <-timer.Chan():
				So("should not be here", ShouldBeTrue)
			default:
			}
			So(timer.Stop(), ShouldBeFalse)
		}
	})

	Convey("TimerReset", t, func() {
		for i := 0; i < 50; i++ {
			ch := make(chan time.Time)
			beg := Now()
			timer := NewTimer(10 * time.Second)

			So(WaitSched(1, time.Second), ShouldBeTrue)

			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			So(timer.Reset(delay), ShouldBeTrue)

			So(WaitSched(1, time.Second), ShouldBeTrue)

			go func() {
				v := <-timer.Chan()
				ch <- v
			}()

			Advance(50 * time.Millisecond)
			runtime.Gosched()
			select {
			case <-ch:
				So("should not be here", ShouldBeNil)
			default:
			}
			Advance(delay - 50*time.Millisecond)

			v := <-ch
			So(v, ShouldEqual, Now())
			So(Now().Sub(beg), ShouldEqual, delay)

			// the timer must has already expired or been stopped.
			So(timer.Stop(), ShouldBeFalse)
		}
	})

	Convey("NewTicker", t, func() {
		for i := 0; i < 50; i++ {
			beg := Now()
			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			ticker := NewTicker(delay)
			So(WaitSched(1, time.Second), ShouldBeTrue)

			go func() { Advance(delay) }()

			end := <-ticker.Chan()
			v := end.Sub(beg)
			So(v, ShouldEqual, delay)

			go func() { Advance(delay) }()

			end = <-ticker.Chan()
			v = end.Sub(beg)
			So(v, ShouldEqual, delay*2)

			ticker.Stop()
			time.Sleep(time.Millisecond)
			select {
			case <-ticker.Chan():
				So("should not be here", ShouldBeTrue)
			default:
			}
		}
	})

	Convey("NewTicker.Zero", t, func() {
		So(func() { NewTicker(0) }, ShouldPanic)
	})

	Convey("Tick", t, func() {
		for i := 0; i < 50; i++ {
			beg := Now()
			delay := time.Duration(rand.Intn(100)+51) * time.Millisecond
			ch := Tick(delay)
			So(WaitSched(1, time.Second), ShouldBeTrue)

			go func() { Advance(delay) }()

			end := <-ch
			v := end.Sub(beg)
			So(v, ShouldEqual, delay)

			go func() { Advance(delay) }()

			end = <-ch
			v = end.Sub(beg)
			So(v, ShouldEqual, delay*2)

		}
	})

	Convey("Tick.Zero", t, func() {
		So(Tick(0), ShouldBeNil)
	})

	Convey("Heap", t, func() {
		n := Now()
		ts := &frozenTimers{}
		for i := 0; i < 50; i++ {
			heap.Push(ts, &frozenTimer{when: n.Add(time.Duration(rand.Int63()))})
		}
		prev := (*ts)[0]
		for len(*ts) != 0 {
			next := heap.Pop(ts).(*frozenTimer)
			So(next.when.Unix(), ShouldBeGreaterThanOrEqualTo, prev.when.Unix())
			prev = next
		}
	})

	Convey("Ticks", t, func() {
		n, m := 3, 100
		ch := make([]<-chan time.Time, n+1)
		for i := 1; i <= n; i++ {
			ch[i] = Tick(time.Duration(i) * time.Millisecond)
		}

		for i := 1; i <= m; i++ {
			Advance(time.Millisecond)
			for j := 1; j <= n; j++ {
				if (i/j)*j == i { // divisible
					select {
					case <-ch[j]:
					default:
						So("should not be here", ShouldBeNil)
					}
				}
			}
		}
	})

}
