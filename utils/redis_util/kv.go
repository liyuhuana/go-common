package redis_util

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//设置过期时间
func Expire(conn redis.Conn, key interface{}, second int32) (err error) {
	_, err = conn.Do(RcExpire, key, second)
	return
}

func Del(conn redis.Conn, key interface{}) (err error) {
	_, err = conn.Do(RcDel, key)
	return
}

//获取到期时间
func Ttl(conn redis.Conn, key interface{}) (ttl int64, err error) {
	ttl, err = redis.Int64(conn.Do(RcTtl, key))
	return
}

func SetKv(conn redis.Conn, key interface{}, value interface{}) (err error) {
	fmt.Println("key", key, "value", value)
	_, err = conn.Do(RcSet, key, value)
	return
}

func GetKv(conn redis.Conn, key interface{}) (value string, err error) {
	return redis.String(conn.Do(RcGet, key))
}

func ExistKv(conn redis.Conn, key interface{}) (exist bool, err error) {
	v, err := redis.Int(conn.Do(RcExists, key))
	if err != nil {
		return
	}
	if v == 1 {    //返回1表示存在
		exist = true
	}

	return
}
