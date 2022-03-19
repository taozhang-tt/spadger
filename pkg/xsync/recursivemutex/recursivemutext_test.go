package recursivemutex

import (
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestRecursiveMutex(t *testing.T) {
	var wg sync.WaitGroup
	var mu RecursiveMutex
	var add1 func(level int)
	var add2 func(level int)
	count := 0

	add1 = func(level int) {
		time.Sleep(time.Millisecond)
		if level > 100 {
			wg.Done()
			return
		}
		mu.Lock()
		if mu.recursion != int64(level) {
			panic(fmt.Sprintf("recursion(%v) not not equal level(%v)\n", mu.recursion, level))
		}
		count += 1
		add1(level + 1)
		mu.Unlock()
	}

	add2 = func(level int) {
		time.Sleep(time.Millisecond)
		if level > 100 {
			wg.Done()
			return
		}
		mu.Lock()
		if mu.recursion != int64(level) {
			panic(fmt.Sprintf("recursion(%v) not not equal level(%v)\n", mu.recursion, level))
		}
		count += 2
		add2(level + 1)
		mu.Unlock()
	}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go add1(1)
		go add2(1)
	}
	wg.Wait()
	if count != 3000 {
		panic(fmt.Sprintf("count(%v) not not equal 3000", count))
	}
}

func TestRecursiveMutexPanic(t *testing.T) {
	Convey("Test panic", t, func() {
		var mu RecursiveMutex
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
		}()
		So(func() { mu.Unlock() }, ShouldPanic)
	})
}
