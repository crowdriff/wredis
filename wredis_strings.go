package wredis

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

// Get fetches a key's string value.
// See http://redis.io/commands/get
func (w *Wredis) Get(key string) (string, error) {
	if key == "" {
		return stringError("key cannot be an empty string")
	}
	var get = func(conn redis.Conn) (string, error) {
		return redis.String(conn.Do("GET", key))
	}
	return w.ExecString(get)
}

// MGet returns the values of all provided keys. For a key that does not exist,
// an empty string is returned. See http://redis.io/commands/mget.
func (w *Wredis) MGet(keys []string) ([]string, error) {
	for _, key := range keys {
		if key == "" {
			return stringsError("keys cannot be empty")
		}
	}
	return w.ExecStrings(func(conn redis.Conn) ([]string, error) {
		args := redis.Args{}.AddFlat(keys)
		return redis.Strings(conn.Do("MGET", args...))
	})
}

// Incr increments the number stored at key by one.
// See http://redis.io/commands/incr
func (w *Wredis) Incr(key string) (int64, error) {
	if key == "" {
		return 0, errors.New("key cannot be an empty string")
	}
	var incr = func(conn redis.Conn) (int64, error) {
		return redis.Int64(conn.Do("INCR", key))
	}
	return w.ExecInt64(incr)
}

// Set sets a key's string value.
// See http://redis.io/commands/set
func (w *Wredis) Set(key, value string) error {
	if key == "" {
		return errors.New("key cannot be an empty string")
	}
	var set = func(conn redis.Conn) (string, error) {
		return redis.String(conn.Do("SET", key, value))
	}
	res, err := w.ExecString(set)
	return checkSimpleStringResponse("Set", res, err)
}

// SetEx sets key's to value with an expiry time measured in seconds.
// See http://redis.io/commands/setex
func (w *Wredis) SetEx(key, value string, seconds int) error {
	if key == "" {
		return errors.New("key cannot be an empty string")
	}
	if seconds <= 0 {
		return errors.New("seconds must be a postive integer")
	}

	var setEx = func(conn redis.Conn) (string, error) {
		args := redis.Args{}.Add(key).Add(seconds).Add(value)
		return redis.String(conn.Do("SETEX", args...))
	}
	res, err := w.ExecString(setEx)
	return checkSimpleStringResponse("SetEx", res, err)
}
