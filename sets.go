package wredis

import "github.com/garyburd/redigo/redis"

// SAdd implements the SADD command. It adds the `members` into the
// set designated by `dest`. An error is returned if `members` is nil
// or empty, otherwise it returns the number of members added.
// See http://redis.io/commands/sadd
func (r *Wredis) SAdd(dest string, members []string) (int64, error) {
	if members == nil || len(members) == 0 {
		return int64Error("cannot SADD empty slice")
	}

	var sadd = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).AddFlat(members)
		return redis.Int64(conn.Do("SADD", args...))
	}
	return r.ExecInt64(sadd)
}

// SCard returns the count of the number of members
// in the set `key`. See http://redis.io/commands/scard
func (r *Wredis) SCard(key string) (int64, error) {
	if key == "" {
		return int64Error("cannot SCARD for an empty key")
	}

	var scard = func(conn redis.Conn) (int64, error) {
		return redis.Int64(conn.Do("SCARD", key))
	}
	return r.ExecInt64(scard)
}

// SDiffStore implements the SDIFFSTORE command. It stores into `dest`
// the set difference of `setA`, `setB`, and `sets` in order.
// See `http://redis.io/commands/sdiffstore`
func (r *Wredis) SDiffStore(dest, setA, setB string, sets ...string) (int64, error) {
	if dest == "" || setA == "" || setB == "" {
		return int64Error("cannot SDIFFSTORE with empty dest or sets")
	}

	var sdiffstore = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).Add(setA).Add(setB).AddFlat(sets)
		return redis.Int64(conn.Do("SDIFFSTORE", args...))
	}
	return r.ExecInt64(sdiffstore)
}

// SMembers returns the members of the set denoted by `key`.
// See http://redis.io/commands/smembers
func (r *Wredis) SMembers(key string) ([]string, error) {
	if key == "" {
		return stringsError("cannot call SMEMBERS for an empty key")
	}

	var smembers = func(conn redis.Conn) ([]string, error) {
		return redis.Strings(conn.Do("SMEMBERS", key))
	}
	return r.ExecStrings(smembers)
}

// SUnionStore implements the SUNIONSTORE command. It stores into
// `dest` the set union of `setA`, `setB`, and `sets` in order.
// See `http://redis.io/commands/sunionstore`
func (r *Wredis) SUnionStore(dest, setA, setB string, sets ...string) (int64, error) {
	if dest == "" || setA == "" || setB == "" {
		return int64Error("cannot SUNIONSTORE with empty dest or sets")
	}

	var sunionstore = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).Add(setA).Add(setB).AddFlat(sets)
		return redis.Int64(conn.Do("SUNIONSTORE", args...))
	}
	return r.ExecInt64(sunionstore)
}
