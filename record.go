package cache

import (
	"sync"
	"time"
)

// record is an entry in the cache with a cached value
// and a time when that value expires.
//
// record is goroutine safe
type record struct {
	sync.RWMutex
	expiresOnUtc time.Time
	value        interface{}
}

// newRecord creates a new record value with the expiration set to Now().UTC().Add(duration)
func newRecord(duration time.Duration, value interface{}) *record {
	var cr record
	cr.update(duration, value)
	return &cr
}

// update mutates the record such that the new expiration time is set to
// Utc() + duration, and the cached value is set to the value parameter. This func should
// only be called from from another record func.
func (cr *record) update(duration time.Duration, value interface{}) interface{} {
	cr.Lock()
	defer cr.Unlock()
	cr.expiresOnUtc = time.Now().UTC().Add(duration)
	cr.value = value
	return value
}

// Value attempts to read the current value of the record. If the value is expired
// genFn is called and the expiration date is updated. The latest value held by
// the record is then returned.
func (cr *record) Value(duration time.Duration, genFn GenFn) (interface{}, error) {
	cr.RLock()
	expiresOnUtc := cr.expiresOnUtc
	value := cr.value
	cr.RUnlock()

	if expiresOnUtc.Before(time.Now().UTC()) {
		newValue, err := genFn()

		if err != nil {
			return value, err
		}

		return cr.update(duration, newValue), nil
	}

	return value, nil
}
