package xtime

import (
	"container/heap"
	"sync"
	"time"
)

type frozenTime struct {
	mu          sync.Mutex
	t           time.Time     // 冻结的时间
	timers      *frozenTimers // 冻结的任务
	waiters     int
	waitersDone chan struct{}
}

func (t *frozenTime) Now() time.Time {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.t
}

func (t *frozenTime) Sleep(d time.Duration) {
	<-t.NewTimer(d).Chan()
}

func (t *frozenTime) NewTimer(d time.Duration) Timer {
	return t.AfterFunc(d, nil)
}

func (t *frozenTime) After(d time.Duration) <-chan time.Time {
	return t.AfterFunc(d, nil).Chan()
}

func (t *frozenTime) AfterFunc(d time.Duration, f func()) Timer {
	timer := &frozenTimer{
		t:    t,
		c:    make(chan time.Time, 1),
		when: t.Now().Add(d),
		f:    f,
	}
	t.startTimer(timer)
	return timer
}

func (t *frozenTime) startTimer(timer *frozenTimer) {
	t.mu.Lock()
	defer t.mu.Unlock()
	timer.stoped = false
	heap.Push(t.timers, timer)

	if t.waiters != 0 && t.waiters <= len(*t.timers) {
		close(t.waitersDone)
	}
}

func (t *frozenTime) stopTimer(timer *frozenTimer) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	if timer.stoped {
		return false
	}

	for idx, val := range *t.timers {
		if timer == val {
			timer.stoped = true
			heap.Remove(t.timers, idx)
			return true
		}
	}
	return false
}

func (t *frozenTime) advance(d time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.t = t.t.Add(d)

	for timer := t.next(); timer != nil; timer = t.next() {
		if timer.c != nil {
			select {
			case timer.c <- timer.when: // 非阻塞
			default:
			}
		}
		if timer.interval != 0 { // 循环类型
			timer.when = timer.when.Add(timer.interval)
			heap.Push(t.timers, timer)
		}
		if timer.f != nil {
			go timer.f()
		}
	}
}

func (t *frozenTime) next() *frozenTimer {
	if len(*t.timers) == 0 {
		return nil
	}
	timer := t.timers.Top()
	if t.t.Before(timer.when) {
		return nil
	}
	heap.Pop(t.timers)
	return timer
}

func (t *frozenTime) waitSched(n int, d time.Duration) bool {
	t.mu.Lock()
	if len(*t.timers) >= n {
		t.mu.Unlock()
		return true
	}
	if t.waiters != 0 {
		panic("cocurrent wait")
	}
	t.waiters = n
	t.waitersDone = make(chan struct{})
	t.mu.Unlock()

	ok := false
	select {
	case <-t.waitersDone:
		ok = true
	case <-time.After(d):
	}
	t.mu.Lock()
	t.waiters = 0
	t.waitersDone = nil
	t.mu.Unlock()
	return ok
}

type frozenTicker struct {
	t *frozenTimer
}

func (t *frozenTicker) Chan() <-chan time.Time {
	return t.t.Chan()
}

func (t *frozenTicker) Stop() {
	t.t.Stop()
}

func (t *frozenTime) NewTicker(d time.Duration) Ticker {
	if d <= 0 {
		panic("non-positive interval for NewTicker")
	}
	timer := &frozenTimer{
		t:        t,
		c:        make(chan time.Time, 1),
		when:     t.Now().Add(d),
		interval: d,
	}
	t.startTimer(timer)
	return &frozenTicker{timer}
}

func (t *frozenTime) Tick(d time.Duration) <-chan time.Time {
	if d <= 0 {
		return nil
	}
	return t.NewTicker(d).Chan()
}

type frozenTimers []*frozenTimer

func (h frozenTimers) Top() *frozenTimer   { return h[0] }
func (h frozenTimers) Len() int            { return len(h) }
func (h frozenTimers) Less(i, j int) bool  { return h[i].when.Before(h[j].when) }
func (h frozenTimers) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *frozenTimers) Push(x interface{}) { *h = append(*h, x.(*frozenTimer)) }
func (h *frozenTimers) Pop() interface{}   { n := len(*h); r := (*h)[n-1]; *h = (*h)[0 : n-1]; return r }

type frozenTimer struct {
	t        *frozenTime
	when     time.Time      // 到期时间
	c        chan time.Time // 到期后发送到期时间
	f        func()         // 到期时要执行的函数
	interval time.Duration  // 循环间隔
	stoped   bool
}

func (t *frozenTimer) Chan() <-chan time.Time {
	return t.c
}

func (t *frozenTimer) Stop() bool {
	return t.t.stopTimer(t)
}

func (t *frozenTimer) Reset(d time.Duration) bool {
	active := t.t.stopTimer(t)
	t.when = t.t.Now().Add(d)
	t.t.startTimer(t)
	return active
}

func (t *frozenTime) init() *frozenTime {
	t.timers = &frozenTimers{}
	return t
}

func (t *frozenTime) freeze(n time.Time) *frozenTime {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.t = n
	return t
}
