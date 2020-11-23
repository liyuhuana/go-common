package redis_util

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//生产者
type Consumer struct {
	Conn  redis.Conn
	Topic []string
}

type Msg struct {
	Topic string
	Value []byte
	Err   error
}

func NewConsumer(conn redis.Conn, topics ...string) (consumer *Consumer, err error) {
	if len(topics) == 0 {
		err = fmt.Errorf("topic is empty")
		return
	}
	consumer = &Consumer{
		Conn: conn,
	}
	for _, topic := range topics {
		consumer.Topic = append(consumer.Topic, topic)
	}

	return
}

func (p *Consumer) Consume() (msg Msg) {

	msg.Value = make([]byte, 0)
	//左出(阻塞)
	args := []interface{}{}
	for _, t := range p.Topic {
		args = append(args, t)
	}
	args = append(args, 0)
	ret, err := redis.ByteSlices(p.Conn.Do(RcBlpop, args...))
	if err != nil {
		msg.Err = err
		return
	}

	if len(ret) != 2 {
		return
	}

	msg.Topic = string(ret[0])
	msg.Value = ret[1]

	return
}

func (p *Consumer) Close() (err error) {
	return p.Conn.Close()
}
