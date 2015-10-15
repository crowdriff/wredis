package wredis

import "github.com/garyburd/redigo/redis"

// SAdd is a helper function for the SADD command. It adds the
// `members` into the set designated by `dest`. An error is returned
// if `members` is nil or empty. See http://redis.io/commands/sadd
func (r *Wredis) SAdd(dest string, members []string) (int64, error) {
	if members == nil || len(members) == 0 {
		return int64Error("Cannot SADD empty slice")
	}

	var sadd Int64 = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).AddFlat(members)
		return redis.Int64(conn.Do("SADD", args...))
	}
	return r.ExecInt64(sadd)
}

// SCard is a function that returns the count of the number of members
// in the set denoted by `key`. See http://redis.io/commands/scard
func (r *Wredis) SCard(key string) (int64, error) {
	if key == "" {
		return int64Error("Cannot SCARD for an empty key")
	}

	var scard Int64 = func(conn redis.Conn) (int64, error) {
		return redis.Int64(conn.Do("SCARD", key))
	}
	return r.ExecInt64(scard)
}

// SDiffStore is a helper function for the SDIFFSTORE command.
// It stores into `dest` the set difference of `setA`, `setB`,
// and `sets`. See `http://redis.io/commands/sdiffstore`
func (r *Wredis) SDiffStore(dest, setA, setB string, sets ...string) (int64, error) {
	if dest == "" || setA == "" || setB == "" {
		return int64Error("Cannot SDIFFSTORE with empty dest or sets")
	}

	var sdiffstore Int64 = func(conn redis.Conn) (int64, error) {
		args := redis.Args{}.Add(dest).Add(setA).Add(setB).AddFlat(sets)
		return redis.Int64(conn.Do("SDIFFSTORE", args...))
	}
	return r.ExecInt64(sdiffstore)
}

// SMembers returns the members of the set denoted by `key`.
// See http://redis.io/commands/smembers
func (r *Wredis) SMembers(key string) ([]string, error) {
	if key == "" {
		return stringsError("Cannot call SMembers on an empty key")
	}

	var smembers Strings = func(conn redis.Conn) ([]string, error) {
		return redis.Strings(conn.Do("SMEMBERS", key))
	}
	return r.ExecStrings(smembers)
}
