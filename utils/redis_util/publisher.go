package redis_util

import (
	"github.com/gomodule/redigo/redis"
)

//生产者
type Publisher struct {
	Conn redis.Conn
}

func NewPublisher(conn redis.Conn) *Publisher {
	return &Publisher{
		Conn: conn,
	}
}

func (p *Publisher) Publish(topic string, message []byte) (err error) {
	_, err = p.Conn.Do(RcPublish, topic, message)
	return
}

func (p *Publisher) Close() (err error) {
	return p.Conn.Close()
}
