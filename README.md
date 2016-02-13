Wredis
===

[![Build Status](https://travis-ci.org/crowdriff/wredis.svg?branch=master)](https://travis-ci.org/crowdriff/wredis) [![Go Report Card](https://goreportcard.com/badge/github.com/crowdriff/wredis)](https://goreportcard.com/report/github.com/crowdriff/wredis)

Wredis is a wrapper around the [redigo](https://github.com/garyburd/redigo) `redis.Pool` and provides an easy-to-use API for [Redis commands](http://redis.io/commands)

## Getting Started

### Go Get

```
go get github.com/crowdriff/wredis
```

### Usage

[API Reference](https://godoc.org/github.com/crowdriff/wredis)

```go
import (
	"log"

	"github.com/crowdriff/wredis"
)

var w *wredis.Wredis

func main() {
	var err error
	if w, err = wredis.NewDefaultPool(); err != nil {
		log.Fatal(err.Error())
	}
	defer w.Close()

	if err = w.Set("mykey", "myval"); err != nil {
		log.Fatal(err.Error())
	}

	val, err := w.Get("mykey")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(val)
}
```

### Implemented Commands
* __Connection__
  * Select: switch the redis database
* __Keys__
  * Del: delete a key
  * Exists: does a key exist
  * Expire: set an expiry time for a key
  * Keys: fetch a list of keys that match the given pattern
  * Rename: rename a key
* __Server__
  * FlushAll: Flush the contents of the redis server (requires unsafe Wredis)
  * FlushDb: Flush the contents of a specific redis db (requires unsafe Wredis)
* __Sets__
  * SAdd: add members to a set
  * SCard: count of a set
  * SDiffStore: perform a diff and store the results in redis
  * SMembers: return the members of a set
  * SUnionStore: perform a union and store the results in redis
* __Strings__
  * Get: get a key's value
  * Incr: increment a key's value by 1
  * Set: set a key's value
  * SetEx: set a key's value with an expiry in seconds

### Convenience methods
* __Keys__
  * DelWithPattern: delete all keys matching a pattern
* __Server__
  * SelectAndFlushDb: selects a db before flushing it
* __Strings__
  * SetExDuration: set a string with an expiry using a `time.Duration`

## Contributing

### Install Tools and Dependencies

```
make tools
make deps
```

### Contribute

1. Fork  
2. Make changes  
3. Add tests  
4. Run `make test`  
5. Send a PR  
