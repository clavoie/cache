package cache

// GenFn represents a func that can generate a cache value.
// Cache values are not updated if an error is returned.
type GenFn func() (interface{}, error)
