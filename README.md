# cache [![GoDoc](https://godoc.org/github.com/clavoie/cache?status.svg)](http://godoc.org/github.com/clavoie/cache) [![Build Status](https://travis-ci.org/clavoie/cache.svg?branch=master)](https://travis-ci.org/clavoie/cache) [![codecov](https://codecov.io/gh/clavoie/cache/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/cache) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/cache)](https://goreportcard.com/report/github.com/clavoie/cache)

An in-memory cache for Go designed to be used with a dependency injection system. The in-memory cache is safe to use across multiple goroutines.

## Usage

The cache implementation is designed to be injected into a wrapper cache structure which handles conversion between `interface{}` and the desired cache value type, as well as determining the cache lifetime.

```go
type ConfigCache struct {
  cache cache.Impl
  duration time.Duration
}

func NewConfigCache(impl cache.Impl) *ConfigCache {
  return &ConfigCache{
    cache: impl,
    duration: time.Minute*5,
  }
}

func (c *ConfigCache) ApiConfig() (*ApiConfig, error) {
  value, err := c.cache.Get("api", c.duration, c.generateApiConfig)
  return value.(*ApiConfig), err
}

func (c *ConfigCache) AdminConfig() (*AdminConfig, error) {
  value, err := c.cache.Get("admin", c.duration, c.generateAdminConfig)
  return value.(*AdminConfig), err
}

func (c *ConfigCache) generateApiConfig() (interface{}, error) { /* db calls, etc */ }
func (c *ConfigCache) generateAdminConfig() (interface{}, error) { /* db calls, etc */ }
```

## Setup

If you are using [di](https://github.com/clavoie/di) the setup can look something like this:

```go
// your_package/di.go
package your_package

import "github.com/clavoie/di"

func NewDiDefs() []*di.Def {
  return []*di.Def{
    // assume ConfigCache is wrapped in an interface at this point
    {NewConfigCache, di.Singleton},
  }
}

// main.go
httpResolver, err = di.NewResolver(onResolveErr,
  your_package.NewDiDefs(),
  cache.NewDiDefs(),
  // etc
)
  
// your_handler.go
func YourHandler(configCache your_package.ConfigCache) {
  config, err := configCache.ApiConfig()
  // etc
}
```
