Wredis
===

[![Build Status](https://travis-ci.org/crowdriff/wredis.svg?branch=master)](https://travis-ci.org/crowdriff/wredis)

Wredis is a wrapper around the [redigo](https://github.com/garyburd/redigo) `redis.Pool` and provides an easy-to-use API for [Redis commands](http://redis.io/commands)

## Getting Started

### Go Get

```
go get github.com/crowdriff/wredis
```

### Usage

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

### Implemented Methods

* __Keys__
  * Del: delete a key
  * Exists: does a key exist
  * Get: get a key's value
  * Rename: rename a key
  * Set: set a key's value
* __Server__
  * FlushAll: Flush the contents of the redis server (requires unsafe Wredis)
* __Sets__
  * SAdd: add members to a set
  * SCard: count of a set
  * SDiffStore: perform a diff and store the results in redis
  * SMembers: return the members of a set
  * SUnionStore: perform a union and store the results in redis

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
