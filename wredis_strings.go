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
