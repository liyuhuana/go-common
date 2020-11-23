package redis_util

import (
	"testing"
	"time"
)

func TestRedisLock_Lock(t *testing.T) {
	rpool := NewPool(addr, pwd, db)
	err := rpool.Get().Err()
	if err != nil {
		t.Error(" error new pool,err:", err)
	}

	rl := RedisLock{
		Key:     "1001",
		TimeOut: 5000,
	}

	lock, err := rl.Lock(rpool.Get())
	if err != nil {
		t.Error("lock err:", err)
		return
	}

	if !lock {
		t.Error("lock fail:", lock)
		return
	}

	//sleep 2s 再次加锁看是否成功
	time.Sleep(time.Millisecond * 2000)

	lock, err = rl.Lock(rpool.Get())
	if err != nil {
		t.Error("lock err:", err)
		return
	}

	if lock {
		t.Error("lock fail:", lock)
		return
	}

	//再 sleep 4s 等锁过期,再加锁看是否成功
	time.Sleep(time.Millisecond * 4000)

	lock, err = rl.Lock(rpool.Get())
	if err != nil {
		t.Error("lock err:", err)
		return
	}

	if !lock {
		t.Error("lock fail:", lock)
		return
	}

}

func TestRedisLock_Unlock(t *testing.T) {
	rpool := NewPool(addr, pwd, db)
	err := rpool.Get().Err()
	if err != nil {
		t.Error(" error new pool,err:", err)
	}

	rl := RedisLock{
		Key:     "1002",
		TimeOut: 500000,
	}

	lock, err := rl.Lock(rpool.Get())
	if err != nil || !lock {
		t.Error("lock err:", err)
		return
	}

	if !lock {
		t.Error("lock fail:", lock)
		return
	}

	//解锁
	err = rl.Unlock(rpool.Get())
	if err != nil {
		t.Error("unlock err:", err)
		return
	}
	//解锁之后再加锁看能否成功
	lock, err = rl.Lock(rpool.Get())
	if err != nil {
		t.Error("lock err:", err)
		return
	}

	if !lock {
		t.Error("lock fail:", lock)
		return
	}

	//清理
	rl.Unlock(rpool.Get())
}
