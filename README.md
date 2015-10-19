Wredis (v0.0.3)
===

Wredis is a wrapper around the __redigo__ `redis.Pool` that provides functions for [Redis commands](http://redis.io/commands)

### Dependencies

`make deps` uses Glock to fetch and manage dependencies

* [redigo](https://github.com/garyburd/redigo)

### Tooling

`make tools` installs the following tools

* [ginkgo](https://github.com/onsi/ginkgo/ginkgo)
* [gomega](https://github.com/onsi/gomega)
* [glock](https://github.com/robfig/glock)
* [golint](https://github.com/golang/lint/golint)

#### Methods

* __Keys__
  * Del: delete a key
  * Exists: does a key exist
  * Rename: rename a key
* __Server__
  * FlushAll: Flush the contents of the redis server (requires unsafe Wredis)
* __Sets__
  * SAdd: add members to a set
  * SCard: count of a set
  * SDiffStore: perform a diff and store the results in redis
  * SMembers: return the members of a set
  * SUnionStore: perform a union and store the results in redis
