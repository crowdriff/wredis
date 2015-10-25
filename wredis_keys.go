package wredis

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// Del deletes one or more keys from Redis. Returns a count of how many
// keys were actually deleted.
// See http://redis.io/commands/del
func (w *Wredis) Del(keys ...string) (int64, error) {
	var del = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.AddFlat(keys)
		return redis.Int64(conn.Do("DEL", args...))
	}
	return w.ExecInt64(del)
}

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

// Exists checks for the existance of `key` in Redis. Note however,
// even though a variable number of keys can be passed to the DEL command
// since Redis 3.0.3, we will restrict this to a single key in order to
// be able to return an absolute response regarding existence.
// See http://redis.io/commands/exists
func (w *Wredis) Exists(key string) (bool, error) {
	if key == "" {
		return boolError("key cannot be empty")
	}

	var exists = func(conn redis.Conn) (int64, error) {
		return redis.Int64(conn.Do("EXISTS", key))
	}

	res, err := w.ExecInt64(exists)
	if err != nil {
		return false, err
	}
	return res == int64(1), nil
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
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("RENAME returned non OK response: %s", res)
	}
	return nil
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
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("SET returned non OK response: %s", res)
	}
	return nil
}
