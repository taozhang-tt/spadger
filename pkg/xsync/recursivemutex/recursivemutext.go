package recursivemutex

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

type RecursiveMutex struct {
	sync.Mutex
	recursion int64 // 锁重入次数
	owner     int64 // 持有锁的 goroutine id
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid { // 锁重入
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	m.recursion--
	if m.recursion == 0 {
		atomic.StoreInt64(&m.owner, -1)
		m.Mutex.Unlock()
	}
}
