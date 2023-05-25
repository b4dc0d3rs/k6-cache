import k6cache from 'k6/x/k6cache';
import { sleep } from 'k6';

export default function () {
  k6cache.createCacheWithExpiryInSeconds(1)
  k6cache.putToDefaultCache('key', 'value')

  k6cache.createWithExpiryInSeconds('token', 1)
  k6cache.putToCache("token", "myToken", "value")

  console.log(`Expecting default key-value to expire, should be 'value', was: ${k6cache.getFromDefaultCache('key')}`)
  console.log(`Expecting named key-value to expire, should be 'value', was: ${k6cache.getFromCache('token', "myToken")}`)

  sleep(2)
  console.log(`Expecting default key-value to expire, should be null: ${k6cache.getFromDefaultCache('key')}`)
  console.log(`Expecting named key-value to expire, should be null: ${k6cache.getFromCache('token', "myToken")}`)
  // true-false tests finished

  try {
    k6cache.getFromCache("invalidCacheName", "irrelevantKey")
  } catch (error) {
    console.log(error)
  }
  try {
    k6cache.putToCache("invalidCacheName", "irrelevantKey")
  } catch (error) {
    console.log(error)
  }

  try {
    k6cache.removeFromCache("invalidCacheName", "irrelevantKey")
  } catch (error) {
    console.log(error)
  }

}