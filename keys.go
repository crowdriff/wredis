package wredis

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// Del deletes a key or keys from redis. The response is
// the number of the keys that were deleted.
// See http://redis.io/commands/del
func (r *Wredis) Del(key string, keys ...string) (int64, error) {
	if key == "" {
		return int64Error("cannot DEL an empty key")
	}
	var otherKeys = []string{key}
	for _, k := range keys {
		if k != "" && key != k {
			otherKeys = append(otherKeys, k)
		}
	}

	var del = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(key).AddFlat(otherKeys)
		return redis.Int64(conn.Do("DEL", args...))
	}
	return r.ExecInt64(del)
}

// Exists checks for the existance of `key` in Redis. Note however,
// even though a variable number of keys can be passed to the DEL command
// since Redis 3.0.3, we will restrict this to a single key in order to
// be able to return an absolute response regarding existence.
// See http://redis.io/commands/exists
func (r *Wredis) Exists(key string) (bool, error) {
	if key == "" {
		return boolError("cannot check EXISTS for an empty key")
	}

	var exists = func(conn redis.Conn) (int64, error) {
		return redis.Int64(conn.Do("EXISTS", key))
	}

	res, err := r.ExecInt64(exists)
	if err != nil {
		return false, err
	}
	return res == int64(1), nil
}

// Rename will rename `key` to `newKey`. They must not be empty
// or identical.
// See `http://redis.io/commands/rename`
func (r *Wredis) Rename(key, newKey string) error {
	if key == "" || newKey == "" {
		return errors.New("cannot RENAME with empty keys")
	}
	if key == newKey {
		return errors.New("cannot RENAME with identical keys")
	}

	var rename = func(conn redis.Conn) (string, error) {
		args := redis.Args{}.Add(key).Add(newKey)
		return redis.String(conn.Do("RENAME", args...))
	}
	res, err := r.ExecString(rename)
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("RENAME returned non OK repsonse: %s", res)
	}
	return nil
}
