package k6cache

import (
	"github.com/patrickmn/go-cache"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/k6cache", new(K6Cache))
}

type K6Cache struct {
	defaultCache *cache.Cache
	caches       map[string]*cache.Cache
}
