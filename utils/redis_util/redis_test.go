package redis_util

import "testing"

var (
	addr = "172.168.163.100:6379"
	pwd  = ""
	db   = 0
)

func TestNewPool(t *testing.T) {
	rpool := NewPool(addr, pwd, db)

	err := rpool.Get().Err()
	if err != nil {
		t.Error(" error new pool,err:", err)
	}
}
