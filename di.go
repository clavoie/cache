package cache

import "github.com/clavoie/di"

// NewDiDefs returns the dependency injection definitions for this package.
//
// The implementation is reoslved per-dependency, so it is recommended to
// wrap Impl in another struct with a longer lifetime.
func NewDiDefs() []*di.Def {
	return []*di.Def{
		&di.Def{NewImpl, di.PerDependency},
	}
}
