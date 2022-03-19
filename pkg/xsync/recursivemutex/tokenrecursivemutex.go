package recursivemutex

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int64
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("%v is not owner token", token))
	}
	m.recursion--
	if m.recursion == 0 {
		atomic.StoreInt64(&m.token, -1)
		m.Mutex.Unlock()
	}
}
