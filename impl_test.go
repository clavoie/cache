package cache_test

import (
	"errors"
	"testing"
	"time"

	"github.com/clavoie/cache"
)

func TestImpl(t *testing.T) {
	var genErr error
	var val int
	cacheDuration := time.Second
	key := "key"

	genFn := func() (interface{}, error) {
		if genErr != nil {
			return nil, genErr
		}

		val += 1
		return val, nil
	}

	c := cache.NewImpl()
	valSnapshot := val

	cacheVal, err := c.Get(key, cacheDuration, genFn)

	if err != nil {
		t.Fatal(err)
	}

	if valSnapshot+1 != cacheVal.(int) {
		t.Fatal(valSnapshot, cacheVal)
	}

	valSnapshot = cacheVal.(int)
	time.Sleep(cacheDuration / 2)

	cacheVal, err = c.Get(key, cacheDuration, genFn)

	if err != nil {
		t.Fatal(err)
	}

	if cacheVal.(int) != valSnapshot {
		t.Fatal(cacheVal, valSnapshot)
	}

	time.Sleep(cacheDuration)

	cacheVal, err = c.Get(key, cacheDuration, genFn)

	if err != nil {
		t.Fatal(err)
	}

	if cacheVal.(int) != valSnapshot+1 {
		t.Fatal(cacheVal, valSnapshot)
	}

	time.Sleep(cacheDuration)
	genErr = errors.New("gen error")

	_, err = c.Get(key, cacheDuration, genFn)

	if err == nil {
		t.Fatal("Was expecting error")
	}

	_, err = c.Get(key+"x", cacheDuration, genFn)

	if err == nil {
		t.Fatal("Was expecting error")
	}
}
