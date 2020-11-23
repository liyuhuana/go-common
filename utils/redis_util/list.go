package redis_util

import (
	"github.com/gomodule/redigo/redis"
)

type List struct {
	ListName string
}

//表头添加数据
func (l List) Lpush(conn redis.Conn, values ...interface{}) (err error) {
	if len(values) == 0 {
		return
	}
	args := []interface{}{l.ListName}
	for _, arg := range values {
		args = append(args, arg)
	}
	_, err = conn.Do(RcLpush, args...)
	return
}

//count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
//count < 0 : 从表尾开始向表头搜索， 移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
//count = 0 : 移除表中所有与 VALUE 相等的值
func (l List) Lrem(conn redis.Conn, count int, value interface{}) (err error) {
	_, err = conn.Do(RcLrem, l.ListName, count, value)
	return
}

// 获取指定范围内值
// 0 表示列表的第一个元素， 1 表示列表的第二个元素，以此类推。
// 也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素
func (l List) Lrange(conn redis.Conn, start, end int) (values []string, err error) {
	values = make([]string, 0)
	values, err = redis.Strings(conn.Do(RcLrange, l.ListName, start, end))
	return
}

//Ltrim 对一个列表进行修剪(trim)，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
//下标 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。
// 也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func (l List) Ltrim(conn redis.Conn, start, end int) (err error) {
	_, err = conn.Do(RcLtrim, l.ListName, start, end)
	return
}

func (l List) Rpush(conn redis.Conn, values ...interface{}) (err error) {
	if len(values) == 0 {
		return
	}
	args := []interface{}{l.ListName}
	for _, arg := range values {
		args = append(args, arg)
	}
	_, err = conn.Do(RcRpush, args...)
	return
}

func (l List) Lpop(conn redis.Conn) (value string, err error) {
	value, err = redis.String(conn.Do(RcLpop, l.ListName))
	return
}

func (l List) Rpop(conn redis.Conn) (value string, err error) {
	value, err = redis.String(conn.Do(RcRpop, l.ListName))
	return
}

func (l List) Llen(conn redis.Conn) (len int64, err error) {
	len, err = redis.Int64(conn.Do(RcLlen, l.ListName))
	return
}
