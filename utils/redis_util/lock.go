package redis_util

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//基于单节点的redis分布式锁
type RedisLock struct {
	Key     string
	TimeOut int64 //ms级别
}

func (p RedisLock) Lock(conn redis.Conn) (success bool, err error) {

	ret, err := String(conn.Do(RcSet, p.key(), p.Key, "Px", p.TimeOut, "Nx"))
	if err != nil && err != ErrNil {
		return
	}
	if err == ErrNil {
		err = nil
	}
	if ret == "OK" {
		success = true
	}
	return
}

func (p RedisLock) Unlock(conn redis.Conn) (err error) {
	_, err = conn.Do(RcDel, p.key())
	return
}

func (p RedisLock) key() string {
	return fmt.Sprintf("__lock_%s__", p.Key)
}
