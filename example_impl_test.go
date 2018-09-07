package cache_test

import (
	"time"

	"github.com/clavoie/cache"
)

type Address struct {
	City   string
	Street string
}

type AddressCache struct {
	cache         cache.Impl
	cacheDuration time.Duration
}

func NewCache() *AddressCache {
	return &AddressCache{
		cache:         cache.NewImpl(),
		cacheDuration: time.Minute * 5,
	}
}

func (ac *AddressCache) Get(key string) (*Address, error) {
	val, err := ac.cache.Get(key, ac.cacheDuration, ac.generateAddress)
	return val.(*Address), err
}

func (ac *AddressCache) generateAddress() (interface{}, error) {
	// db call or something
	return new(Address), nil
}

func ExampleImpl() {

}
