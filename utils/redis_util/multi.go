package redis_util

import "github.com/gomodule/redigo/redis"

func Mulit(conn redis.Conn) (err error) {
	_, err = conn.Do(RcMulti)
	return
}

func DisCard(conn redis.Conn) (err error) {
	_, err = conn.Do(RcDisCard)
	return
}

func Exec(conn redis.Conn) (err error) {
	_, err = conn.Do(RcExec)
	return
}
