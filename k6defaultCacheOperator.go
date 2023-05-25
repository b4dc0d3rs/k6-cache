package k6cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func (k6cache *K6Cache) CreateCacheWithExpiryInSeconds(durationInSeconds int) {
	duration := time.Duration(durationInSeconds) * time.Second
	k6cache.defaultCache = cache.New(duration, duration)
}

func (k6cache *K6Cache) PutToDefaultCache(key string, value string) {
	k6cache.defaultCache.Set(key, value, cache.DefaultExpiration)
}

func (k6cache *K6Cache) GetFromDefaultCache(key string) interface{} {
	stringValue, _ := k6cache.defaultCache.Get(key)
	return stringValue
}

func (k6cache *K6Cache) RemoveFromDefaultCache(key string) {
	k6cache.defaultCache.Delete(key)
}
