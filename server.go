package wredis

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// FlushAll deletes all the keys from all the db's on the Redis
// server. This is a very dangerous method to use in production.
func (r *Wredis) FlushAll() error {
	if r.safe {
		return errors.New("Cannot use FlushAll in safe mode")
	}

	var flushall = func(conn redis.Conn) (string, error) {
		return redis.String(conn.Do("FLUSHALL"))
	}
	res, err := r.ExecString(flushall)
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("FlushAll did not get OK response: %s", res)
	}
	return nil
}
