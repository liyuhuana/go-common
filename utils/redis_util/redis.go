package redis_util

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var (
	ErrNil = redis.ErrNil
)

func NewPool(addr, pwd string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,
		MaxActive:   100,
		Wait:        true,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Println("[ERROR] failed to dial redis server,err:", err)
				return nil, err
			}

			if len(pwd) > 0 {
				_, err = conn.Do(RcAuth, pwd)
				if err != nil {
					log.Println("[ERROR] wrong password ,err:", err)
					conn.Close()
					return nil, err
				}
			}

			_, err = conn.Do(RcSelect, db)
			if err != nil {
				log.Println("[ERROR] failed to select db ,err:", err)
				conn.Close()
				return nil, err
			}
			return conn, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do(RcPing)
			return err
		},
	}
}

func String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

func StringMap(result interface{}, err error) (map[string]string, error) {
	return redis.StringMap(result, err)
}

func Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

func Int64s(reply interface{}, err error) ([]int64, error) {
	return redis.Int64s(reply, err)
}

func Int64Map(reply interface{}, err error) (map[string]int64, error) {
	return redis.Int64Map(reply, err)
}

func Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

func Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}

