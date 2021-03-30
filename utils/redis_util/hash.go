package redis_util

import (
	"github.com/gomodule/redigo/redis"
)

type Hash struct {
	HashName string
}

func (h Hash) Hset(conn redis.Conn, field interface{}, value interface{}) (err error) {
	_, err = conn.Do(RcHset, h.HashName, field, value)
	return
}

func (h Hash) Hget(conn redis.Conn, field interface{}) (value string, err error) {
	value, err = redis.String(conn.Do(RcHget, h.HashName, field))
	return
}

func (h Hash) Hexists(conn redis.Conn, field interface{}) (exist bool, err error) {
	exist, err = redis.Bool(conn.Do(RcHexists, h.HashName, field))
	return
}

func (h Hash) Hdel(conn redis.Conn, fields ...interface{}) (err error) {

	if len(fields) == 0 {
		return
	}
	args := []interface{}{h.HashName}
	for _, arg := range fields {
		args = append(args, arg)
	}
	_, err = conn.Do(RcHdel, args...)
	return
}

func (h Hash) Hmset(conn redis.Conn, fieldValue map[interface{}]interface{}) (err error) {
	if len(fieldValue) == 0 {
		return
	}
	args := []interface{}{h.HashName}

	for field, value := range fieldValue {
		args = append(args, field, value)
	}
	_, err = conn.Do(RcHmset, args...)
	return
}

func (h Hash) Hmget(conn redis.Conn, fields ...string) (fieldValue map[string]string, err error) {

	fieldValue = make(map[string]string)
	if len(fields) == 0 {
		return
	}
	args := []interface{}{h.HashName}
	fieldMap := make(map[int]string)
	for k, field := range fields {
		args = append(args, field)
		fieldMap[k] = field
	}
	values, err := redis.Strings(conn.Do(RcHmget, args...))
	if err != nil {
		return
	}

	for k, v := range values {
		fieldValue[fieldMap[k]] = v
	}

	return
}

func (h Hash) HsetNx(conn redis.Conn, field, value interface{}) (sucess bool, err error) {
	v, err := redis.Int64(conn.Do(RcHsetNx, h.HashName, field, value))
	//如果返回值为1，则field原来不存在，现在设置成功
	//如果返回值为0，则说明field已经存在了，不修改field的值
	if v == 1 {
		sucess = true
	}
	return
}

func (h Hash) Hincrby(conn redis.Conn, field interface{}, increment int) (int64,error) {
	return redis.Int64(conn.Do(RcHincrby, h.HashName, field, increment))
}

func (h Hash) HincrbyFloat(conn redis.Conn, field interface{}, increment float64) (float64 ,error) {
	return redis.Float64(conn.Do(RcHIncrByFloat, h.HashName, field, increment))
}

func (h Hash) HgetAll(conn redis.Conn) (fieldValue map[string]string, err error) {
	fieldValue = make(map[string]string)
	fieldValue, err = redis.StringMap(conn.Do(RcHgetAll, h.HashName))
	return
}

func (h Hash) Hlen(conn redis.Conn) (count int, err error) {
	count, err = redis.Int(conn.Do(RcHlen, h.HashName))
	return
}
