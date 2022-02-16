package xtime

import "time"

type systemTimer struct {
	t *time.Timer
}

func (s *systemTimer) Chan() <-chan time.Time {
	return s.t.C
}

func (s *systemTimer) Stop() bool {
	return s.t.Stop()
}

func (s *systemTimer) Reset(d time.Duration) bool {
	return s.t.Reset(d)
}

type systemTicker struct {
	t *time.Ticker
}

func (s *systemTicker) Chan() <-chan time.Time {
	return s.t.C
}

func (s *systemTicker) Stop() {
	s.t.Stop()
}

func (s *systemTicker) Reset(d time.Duration) {
	s.t.Reset(d)
}

type systemTime struct{}

func (s *systemTime) Now() time.Time {
	return time.Now()
}

func (s *systemTime) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (s *systemTime) NewTimer(d time.Duration) Timer {
	return &systemTimer{time.NewTimer(d)}
}

func (s *systemTime) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (s *systemTime) AfterFunc(d time.Duration, f func()) Timer {
	return &systemTimer{time.AfterFunc(d, f)}
}

func (s *systemTime) NewTicker(d time.Duration) Ticker {
	return &systemTicker{time.NewTicker(d)}
}

func (s *systemTime) Tick(d time.Duration) <-chan time.Time {
	if d <= 0 {
		return nil
	}
	return time.NewTicker(d).C
}
