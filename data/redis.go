
package data

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "gopkg.in/redis.v3"
  "strings"
  "time"
)

func CheckRedisConn() {
  client := redis.NewClient(&redis.Options{
    Addr: config.RedisURI(),
  })
  _, err := client.Ping().Result()
  if err != nil {
    log.Warn("Redis must be running at the URl specificed in config.yaml")
    util.Check(err)
  }
}

func redisConn() *redis.Client {
  return redis.NewClient(&redis.Options{
    Addr: config.RedisURI(),
  })
}

func Set(key string, value string) {
  conn := redisConn()
  err := conn.Set(key, value, 0).Err()
  util.Check(err)
}

func SetTimeout(key string, value string, timeout time.Duration) {
  conn := redisConn()
  err := conn.Set(key, value, timeout).Err()
  util.Check(err)
}

func Get(key string) (bool, string) {
  conn := redisConn()
  resp, err := conn.Get(key).Result()
  if err != nil && strings.Contains(err.Error(), "WRONGTYPE") {
    return false, ""
  }
  if err != nil && strings.Contains(err.Error(), "nil") {
    return false, ""
  }
  util.Check(err)
  return true, resp
}

func Keys(match string) []string {
  conn := redisConn()
  resp, err := conn.Keys(match).Result()
  util.Check(err)
  return resp
}
