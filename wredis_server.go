package wredis

import "github.com/garyburd/redigo/redis"

// FlushAll deletes all the keys from all the db's on the Redis
// server. This is a very dangerous method to use in production.
// See http://redis.io/commands/flushall
func (w *Wredis) FlushAll() error {
	if w.safe {
		return unsafeError("FlushAll")
	}
	var flushall = func(conn redis.Conn) (string, error) {
		return redis.String(conn.Do("FLUSHALL"))
	}
	res, err := w.ExecString(flushall)
	return checkSimpleStringResponse("FlushAll", res, err)
}

// FlushDb deletes all the keys from the currently selected database
// See http://redis.io/commands/flushdb
func (w *Wredis) FlushDb() error {
	if w.safe {
		return unsafeError("FlushDb")
	}
	var flushdb = func(conn redis.Conn) (string, error) {
		return redis.String(conn.Do("FlUSHDB"))
	}
	res, err := w.ExecString(flushdb)
	return checkSimpleStringResponse("FlushDb", res, err)
}
