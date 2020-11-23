package redis_util

import "github.com/gomodule/redigo/redis"

//生产者
type Producer struct {
	Conn redis.Conn
}

func NewProducer(conn redis.Conn) *Producer {
	return &Producer{
		Conn: conn,
	}
}

func (p *Producer) Produce(topic string, message []byte) (err error) {
	//list 右进
	_, err = p.Conn.Do(RcRpush, topic, message)
	return
}

func (p *Producer) Close() (err error) {
	return p.Conn.Close()
}
