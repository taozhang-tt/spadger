package xtime

import "time"

var (
	system            = new(systemTime)
	provider Provider = system
)

func Freeze() { provider = new(frozenTime).init().freeze(time.Now()) }

func FreezeAt(t time.Time) { provider = new(frozenTime).init().freeze(t) }

func Unfreeze() { provider = system }

func Now() time.Time { return provider.Now() }

func Sleep(d time.Duration) { provider.Sleep(d) }

func NewTimer(d time.Duration) Timer { return provider.NewTimer(d) }

func After(d time.Duration) <-chan time.Time { return provider.After(d) }

func AfterFunc(d time.Duration, f func()) Timer { return provider.AfterFunc(d, f) }

func NewTicker(d time.Duration) Ticker { return provider.NewTicker(d) }

func Tick(d time.Duration) <-chan time.Time { return provider.Tick(d) }

func Advance(d time.Duration) {
	frozen, ok := provider.(*frozenTime)
	if !ok {
		panic("must be freezed")
	}
	frozen.advance(d)
}

func WaitSched(n int, d time.Duration) bool {
	frozen, ok := provider.(*frozenTime)
	if !ok {
		panic("must be freezed")
	}
	return frozen.waitSched(n, d)
}

type (
	Provider interface {
		Now() time.Time
		Sleep(d time.Duration)
		NewTimer(d time.Duration) Timer
		After(d time.Duration) <-chan time.Time
		AfterFunc(d time.Duration, f func()) Timer
		NewTicker(d time.Duration) Ticker
		Tick(d time.Duration) <-chan time.Time
	}

	// Timer see time.Timer
	Timer interface {
		Chan() <-chan time.Time
		Stop() bool
		Reset(d time.Duration) bool
	}

	// Ticker see time.Ticker
	Ticker interface {
		Chan() <-chan time.Time
		Stop()
	}
)
