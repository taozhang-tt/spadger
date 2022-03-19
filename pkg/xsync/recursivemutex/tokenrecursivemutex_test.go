package recursivemutex

import (
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestTokenRecursiveMutex(t *testing.T) {
	var wg sync.WaitGroup
	var mu TokenRecursiveMutex
	var add1 func(level int, token int64)
	var add2 func(level int, token int64)
	count := 0

	add1 = func(level int, token int64) {
		time.Sleep(time.Millisecond)
		if level > 100 {
			wg.Done()
			return
		}
		mu.Lock(token)
		if mu.recursion != int64(level) {
			panic(fmt.Sprintf("recursion(%v) not not equal level(%v)\n", mu.recursion, level))
		}
		count += 1
		add1(level+1, token)
		mu.Unlock(token)
	}

	add2 = func(level int, token int64) {
		time.Sleep(time.Millisecond)
		if level > 100 {
			wg.Done()
			return
		}
		mu.Lock(token)
		if mu.recursion != int64(level) {
			panic(fmt.Sprintf("recursion(%v) not not equal level(%v)\n", mu.recursion, level))
		}
		count += 2
		add2(level+1, token)
		mu.Unlock(token)
	}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go add1(1, int64(i))
		go add2(1, 10*int64(i))
	}
	wg.Wait()
	if count != 3000 {
		panic(fmt.Sprintf("count(%v) not not equal 3000", count))
	}
}

func TestTokenRecursiveMutexPanic(t *testing.T) {
	Convey("Test panic", t, func() {
		var mu TokenRecursiveMutex
		go func() {
			mu.Lock(1)
			time.Sleep(time.Second)
		}()
		So(func() { mu.Unlock(2) }, ShouldPanic)
	})
}
