package wredis

import (
	"errors"
	"fmt"
	"time"

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
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("SET returned non OK response: %s", res)
	}
	return nil
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
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("SETEX returned non OK response: %s", res)
	}
	return nil
}

// SetExDuration is a convenience method to set a key's value with and expiry time.
func (w *Wredis) SetExDuration(key, value string, duration time.Duration) error {
	seconds := int(duration.Seconds())
	if seconds <= 0 {
		return errors.New("duration must be at least 1 second")
	}
	return w.SetEx(key, value, seconds)
}