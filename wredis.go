package wredis

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

const defaultHost = "localhost"
const defaultPort = 6379
const defaultDb = 0

// Wredis is a struct wrapper around the redis.Pool
// that implements Redis commands (http://redis.io/commands)
type Wredis struct {
	pool *redis.Pool
	safe bool
}

// Close closes the pool connection
func (w *Wredis) Close() error {
	return w.pool.Close()
}

// NewDefaultPool returns a redis.Pool with a localhost:6379 address
// and db set to 0.
func NewDefaultPool() (*Wredis, error) {
	return NewPool(defaultHost, defaultPort, defaultDb)
}

// NewPool creates a redis pool connected to the given host:port and db.
func NewPool(host string, port, db uint) (*Wredis, error) {
	if host == "" {
		return nil, errors.New("host cannot be empty")
	}
	if port == 0 {
		return nil, errors.New("port cannot be 0")
	}
	addr := fmt.Sprintf("%s:%d", host, int(port))
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialDatabase(int(db)))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &Wredis{pool, true}, nil
}

// NewUnsafe returns an unsafe wrapper around the redigo Pool.
// The lack of safety allows usage of certain methods that could be
// harmful if accidentally invoked in production (e.g. FlushAll)
func NewUnsafe(host string, port, db uint) (*Wredis, error) {
	w, err := NewPool(host, port, db)
	if err != nil {
		return nil, err
	}
	w.safe = false
	return w, nil
}

// ExecInt64 is a helper function to execute any series of commands
// on a redis.Conn that return an int64 response
func (w *Wredis) ExecInt64(f func(redis.Conn) (int64, error)) (int64, error) {
	conn := w.pool.Get()
	defer conn.Close()
	return f(conn)
}

// ExecString is a helper function to execute any series of commands
// on a redis.Conn that return a string response
func (w *Wredis) ExecString(f func(redis.Conn) (string, error)) (string, error) {
	conn := w.pool.Get()
	defer conn.Close()
	return f(conn)
}

// ExecStrings is a helper function to execute any series of commands
// on a redis.Conn that return a string slice response
func (w *Wredis) ExecStrings(f func(redis.Conn) ([]string, error)) ([]string, error) {
	conn := w.pool.Get()
	defer conn.Close()
	return f(conn)
}
