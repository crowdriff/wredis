package wredis

import "github.com/garyburd/redigo/redis"

// Del deletes a key or keys from redis. See http://redis.io/commands/del
func (r *Wredis) Del(key string, keys ...string) (int64, error) {
	if key == "" {
		return int64Error("Cannot DEL empty key")
	}
	var otherKeys = []string{key}
	for _, k := range keys {
		if k != "" && key != k {
			otherKeys = append(otherKeys, k)
		}
	}

	var del Int64 = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(key).AddFlat(otherKeys)
		return redis.Int64(conn.Do("DEL", args...))
	}
	return r.ExecInt64(del)
}

// Exists checks for the existance of `key` in Redis. Note however,
// that even though Redis 3.0.3 allows a variable number of keys to
// be passed, we will restrict this to a single key in order to be able
// to return an absolute response regarding said existence.
// See http://redis.io/commands/exists
func (r *Wredis) Exists(key string) (bool, error) {
	if key == "" {
		return boolError("Cannot check EXISTS for empty key")
	}

	var exists Int64 = func(conn redis.Conn) (int64, error) {
		return redis.Int64(conn.Do("EXISTS", key))
	}

	res, err := r.ExecInt64(exists)
	if err != nil {
		return false, err
	}
	return res == int64(1), nil
}

// Rename is a helper function to renmae `key` to `newKey`.
// See `http://redis.io/commands/rename`
func (r *Wredis) Rename(key, newKey string) (string, error) {
	if key == "" || newKey == "" {
		return stringError("Cannot RENAME with empty keys")
	}
	if key == newKey {
		return stringError("key and newKey are identical")
	}

	var rename String = func(conn redis.Conn) (string, error) {
		args := redis.Args{}.Add(key).Add(newKey)
		return redis.String(conn.Do("RENAME", args...))
	}
	return r.ExecString(rename)
}
