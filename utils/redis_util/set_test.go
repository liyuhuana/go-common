package redis_util

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func newPool() (rpool *redis.Pool, err error) {
	rpool = NewPool("0.0.0.0:6000", "", 15)
	err = rpool.Get().Err()
	return
}

func TestSet_Sadd_Members(t *testing.T) {

	var (
		setName = "test"
		value   = "zhangsan"
	)
	s := Set{
		SetName: setName,
	}
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()
	err = s.Sadd(coon, value)
	if err != nil {
		t.Error("set Sadd err:", err)
		return
	}

	values, err := s.Smembers(coon)
	if err != nil || len(values) != 1 {
		t.Error("members err:", err, "members:", values)
		return
	}

	if values[0] != value {
		t.Error("set Sadd err:", err, "members:", values)
		return
	}

	Del(coon,s.SetName)
}


func TestSet_RemoveMembers(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		setName     = "test"
		originValue = [3]interface{}{"zhangsan", "lisi", "wangwu"}
		removeValue = originValue[0:2]
		expectValue = originValue[2]
	)
	s := Set{
		SetName: setName,
	}
	for _, value := range originValue {
		err := s.Sadd(coon, value)
		if err != nil {
			t.Error("set Sadd err:", err)
			return
		}
	}

	err = s.Srem(coon, removeValue...)
	if err != nil {
		t.Error("set remove member err:", err)
		return
	}

	members, err := s.Smembers(coon)
	if err != nil || len(members) != 1 || members[0] != expectValue {
		t.Error("remove member err:", err)
	}
	Del(coon,s.SetName)
}

func TestSet_IsMember(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		setName       = "test"
		originValue   = [3]string{"zhangsan", "lisi", "wangwu"}
		existValue    = originValue[1]
		notExistValue = "rheih"
	)
	s := Set{
		SetName: setName,
	}

	for _, value := range originValue {
		err := s.Sadd(coon, value)
		if err != nil {
			t.Error("set Sadd err:", err)
			return
		}
	}

	exist, err := s.SisMember(coon, existValue)
	if err != nil || !exist {
		t.Error("set ismember err:", err)
		return
	}

	exist, err = s.SisMember(coon, notExistValue)
	if err != nil || exist {
		t.Error("set ismember err:", err)
		return
	}
	Del(coon,s.SetName)

}

func TestSet_RandMembers(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		setName     = "test"
		originValue = [3]string{"zhangsan", "lisi", "wangwu"}
		existMap    = map[string]bool{
			originValue[0]: true,
			originValue[1]: true,
			originValue[2]: true,
		}
	)
	s := Set{
		SetName: setName,
	}

	for _, value := range originValue {
		err := s.Sadd(coon, value)
		if err != nil {
			t.Error("set Sadd err:", err)
			return
		}
	}

	members, err := s.SrandMembers(coon, 2)
	t.Log("members:", members)
	if err != nil || len(members) != 2 || !existMap[members[0]] || !existMap[members[1]] {
		t.Error("set randmember err:", err)
		return
	}
	Del(coon,s.SetName)
}
