package wredis

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

// MakeDefaultPool returns a redis.Pool with a localhost:6379 address
// and db set to 0
func MakeDefaultPool() *redis.Pool {
	return MakePool("localhost", "6379", 0)
}

// MakePool creates a redis pool connected to the given host:port and db.
func MakePool(redisHost, redisPort string, db int) *redis.Pool {
	if db < 0 {
		log.Fatal("cannot have a negative db")
	}
	address := redisHost + ":" + redisPort
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialDatabase(db))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
