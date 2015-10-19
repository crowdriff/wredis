package wredis

import (
	"errors"
	"log"

	"github.com/garyburd/redigo/redis"
)

// Wredis is a struct wrapper around the redis.Pool
// that implements Redis commands (http://redis.io/commands)
type Wredis struct {
	pool *redis.Pool
	safe bool
}

// New returns a new Wredis object
func New(pool *redis.Pool) *Wredis {
	if pool == nil {
		log.Fatal("calling New with nil redis.Pool")
	}
	return &Wredis{pool, true}
}

// NewUnsafe returns an unsafe wrapper around the redis.Pool
// that implements the Redis commands `http://redis.io/commands.
// The lack of safety allows usage of certain methods that could be
// harmful if accidentally invoked in production (e.g. FlushAll)
func NewUnsafe(pool *redis.Pool) *Wredis {
	if pool == nil {
		log.Fatal("calling NewUnsafe with nil redis.Pool")
	}
	return &Wredis{pool, false}
}

// Close closes the pool connection
func (r *Wredis) Close() error {
	return r.pool.Close()
}

// ExecInt64 is a helper function to execute any series of commands
// on a redis.Conn that return an int64 response
func (r *Wredis) ExecInt64(f func(redis.Conn) (int64, error)) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return f(conn)
}

// ExecString is a helper function to execute any series of commands
// on a redis.Conn that return a string response
func (r *Wredis) ExecString(f func(redis.Conn) (string, error)) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return f(conn)
}

// ExecStrings is a helper function to execute any series of commands
// on a redis.Conn that return a string slice response
func (r *Wredis) ExecStrings(f func(redis.Conn) ([]string, error)) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return f(conn)
}

//
// error helper functions
//

func boolError(msg string) (bool, error) {
	return false, errors.New(msg)
}

func int64Error(msg string) (int64, error) {
	return int64(0), errors.New(msg)
}

func stringError(msg string) (string, error) {
	return "", errors.New(msg)
}

func stringsError(msg string) ([]string, error) {
	return nil, errors.New(msg)
}
