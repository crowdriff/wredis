package wredis

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

// Select selects the Database specified by the parameter.
// We use an unsigned int because Redis databases are numbered
// using a zero based index.
// See http://redis.io/commands
func (w *Wredis) Select(db int) error {
	if db < 0 {
		return errors.New("db index must be 0 or positive")
	}
	var _select = func(conn redis.Conn) (string, error) {
		args := redis.Args{}.Add(db)
		return redis.String(conn.Do("SELECT", args...))
	}
	res, err := w.ExecString(_select)
	return checkSimpleStringResponse("Select", res, err)
}
