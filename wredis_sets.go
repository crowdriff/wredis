package wredis

import "github.com/garyburd/redigo/redis"

// SAdd implements the SADD command. It adds the `members` into the
// set `dest`. An error is returned if `members` is nil or empty,
// otherwise it returns the number of members added.
// See http://redis.io/commands/sadd
func (w *Wredis) SAdd(dest string, members []string) (int64, error) {
	if members == nil || len(members) == 0 {
		return int64Error("members cannot be an empty slice")
	}
	var sadd = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).AddFlat(members)
		return redis.Int64(conn.Do("SADD", args...))
	}
	return w.ExecInt64(sadd)
}

// SCard returns the cardinality of the set `key`.
// See http://redis.io/commands/scard
func (w *Wredis) SCard(key string) (int64, error) {
	if key == "" {
		return int64Error("key cannot be empty")
	}
	var scard = func(conn redis.Conn) (int64, error) {
		return redis.Int64(conn.Do("SCARD", key))
	}
	return w.ExecInt64(scard)
}

// SDiffStore executes the SDIFFSTORE command.
// See `http://redis.io/commands/sdiffstore`
func (w *Wredis) SDiffStore(dest string, sets ...string) (int64, error) {
	if dest == "" {
		return int64Error("dest cannot be an empty string")
	}
	if len(sets) == 0 {
		return int64Error("SDiffStore requires atleast 1 input set")
	}
	for _, s := range sets {
		if s == "" {
			return int64Error("set keys cannot be empty strings")
		}
	}
	var sdiffstore = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).AddFlat(sets)
		return redis.Int64(conn.Do("SDIFFSTORE", args...))
	}
	return w.ExecInt64(sdiffstore)
}

// SMembers returns the members of the set denoted by `key`.
// See http://redis.io/commands/smembers
func (w *Wredis) SMembers(key string) ([]string, error) {
	if key == "" {
		return stringsError("key cannot be an empty string")
	}
	var smembers = func(conn redis.Conn) ([]string, error) {
		return redis.Strings(conn.Do("SMEMBERS", key))
	}
	return w.ExecStrings(smembers)
}

// SUnionStore implements the SUNIONSTORE command.
// See `http://redis.io/commands/sunionstore`
func (w *Wredis) SUnionStore(dest string, sets ...string) (int64, error) {
	if dest == "" {
		return int64Error("dest cannot be an empty string")
	}
	if len(sets) == 0 {
		return int64Error("SUnionStore requires at least 1 input set")
	}
	for _, s := range sets {
		if s == "" {
			return int64Error("set keys cannot be empty strings")
		}
	}
	var sunionstore = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).AddFlat(sets)
		return redis.Int64(conn.Do("SUNIONSTORE", args...))
	}
	return w.ExecInt64(sunionstore)
}
