# k6cache

A bunch of random functions for k6 performance testing that I found missing, but useful in our work.

# Compile for development
```sh
xk6 build v0.41.0 \
  --with github.com/b4dc0d3rs/k6-cache=.

./k6 run k6cache.js
```

# Use

Import this:
```js
import k6cache from 'k6/x/k6cache'
```

## Default expiring cache
Default cache has automatic expiry measured in seconds. There is one default cache that share the same expiry duration.

Methods that share this cache have `Default` in the name.

### Usage examples:
```js
// Configure it in the setup method
k6utils.createCacheWithExpiryInSeconds(1)

// Insert anything anytime
k6utils.putToDefaultCache('key', 'value')

// Get anywhere anytime, even in a different method
k6utils.getFromDefaultCache('key')

k6utils.removeFromDefaultCache('key')
```

## Named caches
`k6-cache` stores a map of caches. Each cache can have different autoexpiry duration.

Use this when single expiry duration in the default cache is not enough.

To create a named cache with expiry in seconds:
```js
k6cache.createWithExpiryInSeconds('cache_name', 1)
```
Named cache must be created before use. If not, an error will be thrown.

To put key-value to the named cache:
```js
k6cache.putToCache("cache_name", "key", "value")
```

The same pattern applies when inserting and deleting from a named cache:
```js
k6cache.getFromCache("cache_name", "key")

k6cache.removeFromCache("cache_name", "key")
```