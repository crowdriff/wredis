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

// Int64 is a func type that accepts a redis.Conn object and
// returns an int64 response
type Int64 func(redis.Conn) (int64, error)

// ExecInt64 is a helper function to execute any series of commands
// that return an int64 response
func (r *Wredis) ExecInt64(f Int64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return f(conn)
}

// String is a func type that accepts a redis.Conn object and
// returns a string respnose
type String func(redis.Conn) (string, error)

// ExecString is a helper function to execute any series of commands
// that return a string response
func (r *Wredis) ExecString(f String) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return f(conn)
}

// Strings is a func type that accepts a redis.Conn object and
// returns a string slice response
type Strings func(redis.Conn) ([]string, error)

// ExecStrings is a helper function to execute any series of commands
// that return a string slice response
func (r *Wredis) ExecStrings(f Strings) ([]string, error) {
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
