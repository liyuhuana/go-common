package redis_util

import "testing"

func TestSetKv_Get_Del(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	key := "hello"
	value := "world"

	err = SetKv(coon, key, value)
	if err != nil {
		t.Fatal("failed to set key value,err:", err)
	}

	v, err := GetKv(coon, key)
	if err != nil || v != "world" {
		t.Fatal("failed to get key value,err:", err)
	}

	err = Del(coon, key)
	if err != nil {
		t.Fatal("failed to del key,err:", err)
	}

	v, err = GetKv(coon, key)
	if (err != nil && err!= ErrNil) || v != "" {
		t.Fatal("failed to del value,err:", err)
	}
}

func TestExpire_TTL_Exist(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	key := "hello"
	value := "world"

	err = SetKv(coon, key, value)
	if err != nil {
		t.Fatal("failed to set key value,err:", err)
	}

	err= Expire(coon,key,300)
	if err!=nil{
		t.Fatal("failed to expire key,err:",err)
	}

	ttl,err:= Ttl(coon,key)
	if err!=nil || ttl<=0 {
		t.Fatal("failed to ttl key,err:",err,ttl)
	}

	t.Log("key ttl:",ttl)

	//先判断之前是否存在
	exist,err:= ExistKv(coon,key)
	if err!=nil || !exist{
		t.Fatal("failed to exist key,err:",err,exist)
	}

	//删除后在判断是否还存在
	Del(coon,key)
	exist,err= ExistKv(coon,key)
	if err!=nil || exist{
		t.Fatal("failed to exist key,err:",err,exist)
	}
}
