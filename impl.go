package cache

import (
	"sync"
	"time"
)

// Impl represents an in-memory cache of objects.
type Impl interface {
	// Get returns an item from the cache based on its key. If it cannot
	// be found it calls genFn and caches the results for durations.
	Get(key string, duration time.Duration, genFn GenFn) (interface{}, error)
}

// impl is an implementation of Impl
type impl struct {
	sync.RWMutex
	records map[string]*record
}

// NewImpl returns a new instance of Impl.
func NewImpl() Impl {
	return &impl{
		records: make(map[string]*record, 10),
	}
}

func (c *impl) Get(key string, duration time.Duration, genFn GenFn) (interface{}, error) {
	c.RLock()
	record, hasRecord := c.records[key]
	c.RUnlock()

	if hasRecord {
		return record.Value(duration, genFn)
	}

	value, err := genFn()

	if err != nil {
		return nil, err
	}

	record = newRecord(duration, value)
	c.Lock()
	defer c.Unlock()
	c.records[key] = record
	return value, nil
}
