package k6cache

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func (k6cache *K6Cache) CreateWithExpiryInSeconds(name string, durationInSeconds int) {
	if k6cache.caches == nil {
		k6cache.caches = make(map[string]*cache.Cache)
	}
	namedCache := k6cache.caches[name]
	duration := time.Duration(durationInSeconds) * time.Second

	if namedCache == nil {
		namedCache = cache.New(duration, duration)
		k6cache.caches[name] = namedCache
	}
}

func (k6cache *K6Cache) PutToCache(name string, key string, value string) error {
	namedCache := k6cache.caches[name]
	if err := validateCacheExistence(name, namedCache); err != nil {
		return err
	}

	namedCache.Set(key, value, cache.DefaultExpiration)
	return nil
}

func (k6cache *K6Cache) GetFromCache(name string, key string) (interface{}, error) {
	namedCache := k6cache.caches[name]
	if err := validateCacheExistence(name, namedCache); err != nil {
		return nil, err
	}
	stringValue, _ := namedCache.Get(key)
	return stringValue, nil
}

func (k6cache *K6Cache) RemoveFromCache(name string, key string) error {
	namedCache := k6cache.caches[name]
	if err := validateCacheExistence(name, namedCache); err != nil {
		return err
	}
	namedCache.Delete(key)
	return nil
}

func validateCacheExistence(name string, namedCache *cache.Cache) error {
	if namedCache == nil {
		return fmt.Errorf("Cache with name %s was not created, execute `createWithExpiryInSeconds(\"%s\", <expiry>)` first", name, name)
	}
	return nil
}
