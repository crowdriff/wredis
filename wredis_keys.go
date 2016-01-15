package wredis

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

// Del deletes one or more keys from Redis and returns a count of how many
// keys were actually deleted.
// See http://redis.io/commands/del
func (w *Wredis) Del(keys ...string) (int64, error) {
	if keys == nil || len(keys) == 0 {
		return int64Error("must provide at least 1 key")
	}
	for _, key := range keys {
		if "" == key {
			return int64Error("keys cannot be empty strings")
		}
	}
	var del = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.AddFlat(keys)
		return redis.Int64(conn.Do("DEL", args...))
	}
	return w.ExecInt64(del)
}

// Exists checks for the existance of `key` in Redis. Note however,
// even though a variable number of keys can be passed to the EXISTS command
// since Redis 3.0.3, we will restrict this to a single key in order to
// be able to return an absolute response regarding existence.
// See http://redis.io/commands/exists
func (w *Wredis) Exists(key string) (bool, error) {
	if key == "" {
		return boolError("key cannot be empty")
	}
	var exists = func(conn redis.Conn) (bool, error) {
		return redis.Bool(conn.Do("EXISTS", key))
	}
	return w.ExecBool(exists)
}

// Expire sets a timeout of "seconds" on "key".
// If an error is encountered, it is returned. If the key doesn't exist or the
// timeout couldn't be set, `false, nil` is returned. On success, `true, nil`
// is returned.
// See http://redis.io/commands/expire
func (w *Wredis) Expire(key string, seconds int) (bool, error) {
	if key == "" {
		return false, errors.New("key cannot be an empty string")
	}
	var expire = func(conn redis.Conn) (bool, error) {
		return redis.Bool(conn.Do("EXPIRE", key, seconds))
	}
	return w.ExecBool(expire)
}

// Keys takes a pattern and returns any/all keys matching the pattern.
// See http://redis.io/commands/keys
func (w *Wredis) Keys(pattern string) ([]string, error) {
	if pattern == "" {
		return stringsError("pattern cannot be empty")
	}
	var keys = func(conn redis.Conn) ([]string, error) {
		return redis.Strings(conn.Do("KEYS", pattern))
	}
	return w.ExecStrings(keys)
}

// Rename will rename `key` to `newKey`. They must not be empty
// or identical.
// See `http://redis.io/commands/rename`
func (w *Wredis) Rename(key, newKey string) error {
	if key == "" || newKey == "" {
		return errors.New("key and newKey cannot be empty")
	}
	if key == newKey {
		return errors.New("key cannot be equal to newKey")
	}
	var rename = func(conn redis.Conn) (string, error) {
		return redis.String(conn.Do("RENAME", key, newKey))
	}
	res, err := w.ExecString(rename)
	return checkSimpleStringResponse("Rename", res, err)
}
