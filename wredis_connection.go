package wredis

import "github.com/garyburd/redigo/redis"

// Select selects the Database specified by the parameter.
// We use an unsigned int because Redis databases are numbered
// using a zero based index.
// See http://redis.io/commands
func (w *Wredis) Select(db uint) error {
	var _select = func(conn redis.Conn) (string, error) {
		return redis.String(conn.Do("SELECT", db))
	}
	res, err := w.ExecString(_select)
	return checkSimpleStringResponse("Select", res, err)
}
