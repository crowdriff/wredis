package wredis

import (
	"github.com/garyburd/redigo/redis"
)

// LPush inserts the provided item(s) at the head of the list stored at key. For
// more information, see http://redis.io/commands/lpush.
func (w *Wredis) LPush(key string, items ...string) (int64, error) {
	if key == "" {
		return int64Error("key cannot be empty")
	}
	if len(items) == 0 {
		return int64Error("must provide at least one item")
	}
	for _, i := range items {
		if i == "" {
			return int64Error("an item cannot be empty")
		}
	}
	return w.ExecInt64(func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(key).AddFlat(items)
		return redis.Int64(conn.Do("LPUSH", args...))
	})
}
