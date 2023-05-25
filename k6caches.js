import k6cache from 'k6/x/k6cache';
import { sleep } from 'k6';

export default function () {
  k6cache.createCacheWithExpiryInSeconds(1)
  k6cache.putToCache('key', 'value')

  k6cache.createWithExpiryInSeconds('token', 1)
  k6cache.putToNamedCache("token", "myToken", "value")

  console.log(`Expecting default key-value to expire, should be 'value', was: ${k6cache.getFromCache('key')}`)
  console.log(`Expecting named key-value to expire, should be 'value', was: ${k6cache.getFromNamedCache('token', "myToken")}`)

  sleep(2)
  console.log(`Expecting default key-value to expire, should be null: ${k6cache.getFromCache('key')}`)
  console.log(`Expecting named key-value to expire, should be null: ${k6cache.getFromNamedCache('token', "myToken")}`)
  // true-false tests finished

  try {
    k6cache.getFromNamedCache("invalidCacheName", "irrelevantKey")
  } catch (error) {
    console.log(error)
  }
  try {
    k6cache.putToNamedCache("invalidCacheName", "irrelevantKey")
  } catch (error) {
    console.log(error)
  }

  try {
    k6cache.removeFromNamedCache("invalidCacheName", "irrelevantKey")
  } catch (error) {
    console.log(error)
  }

}