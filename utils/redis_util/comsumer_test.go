package redis_util

import "testing"

func TestConsumer_Consume(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		topics = []string{"topic_test1", "topic_test2"}
	)

	producer := NewProducer(coon)
	defer producer.Close()

	comsumer, err := NewConsumer(coon, topics...)
	defer comsumer.Close()
	if err != nil {
		t.Error("failed to new consumer,err:", err)
		return
	}

	err = producer.Produce(topics[0], []byte("hello world"))
	if err != nil {
		t.Error("producer err:", err)
		return
	}

	msg := comsumer.Consume()
	if msg.Err != nil || msg.Topic != topics[0] || string(msg.Value) != "hello world" {
		t.Error("comsume err:", msg)
		return
	}

	t.Log("msg:", msg)

	err = producer.Produce(topics[1], []byte("hello world!"))
	if err != nil {
		t.Error("producer err:", err)
		return
	}

	msg = comsumer.Consume()
	if msg.Err != nil || msg.Topic != topics[1] || string(msg.Value) != "hello world!" {
		t.Error("comsume err:", msg)
		return
	}

	t.Log("msg:", msg)
}
