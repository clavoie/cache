package cache_test

import (
	"testing"

	"github.com/clavoie/cache"
)

func TestNewDiDefs(t *testing.T) {
	defs := cache.NewDiDefs()

	if defs == nil {
		t.Fatal("Expecting non-nil defs")
	}

	defs2 := cache.NewDiDefs()
	if defs[0] == defs2[0] {
		t.Fatal("Not expecting defs to match")
	}
}
