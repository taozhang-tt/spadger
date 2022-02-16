package xtime

import (
	"errors"
	"sync"
	"time"
)

var (
	locs   = map[string]*time.Location{}
	locsMu sync.RWMutex

	errCacheLocation = errors.New("xtime: invalid location from cache")
)

func GetLocation(name string) (loc *time.Location, err error) {
	var ok bool
	locsMu.RLock()
	loc, ok = locs[name]
	locsMu.RUnlock()
	if !ok {
		loc, err = time.LoadLocation(name)
		locsMu.Lock()
		locs[name] = loc
		locsMu.Unlock()
		if err != nil {
			return nil, err
		}
	}
	if loc == nil {
		return loc, errCacheLocation
	}
	return loc, nil
}
