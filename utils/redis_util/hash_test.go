package redis_util

import "testing"

func TestHash_Hset_Hget(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "name"
		value    = "zhangsan"
	)
	h := Hash{
		HashName: hashName,
	}

	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hset err:", err)
		return
	}

	v, err := h.Hget(coon, field)
	if err != nil || v != value {
		t.Error("hash hset_hget err:", err)
	}
	Del(coon,h.HashName)

}

func TestHash_Hmset_Hmget_Hdel(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName   = "test"
		filedValue = map[interface{}]interface{}{
			"name": "zhangsan",
			"age":  17,
			1:  "man",
		}
	)

	h := Hash{
		HashName: hashName,
	}
	err = h.Hmset(coon, filedValue)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	fv, err := h.Hmget(coon, "name", "age", "1")
	if err != nil && len(fv) != 3 && fv["name"] != "zhangsan" && fv["age"] != "17" && fv["1"] != "man" {
		t.Error("hash hget err:", err)
		return
	}

	err1 := h.Hdel(coon, "name")
	fv, err = h.Hmget(coon, "name", "age", "1")
	if err1 != nil || err != nil ||
		len(fv) != 3 || fv["name"] != "" || fv["age"] != "17" || fv["1"] != "man" {
		t.Error("hash hdel err:", err)
		return
	}

	Del(coon,h.HashName)

}

func TestHash_HsetNx_Del(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()
	var (
		hashName = "test"
		field    = "name"
		value    = "zhangsan"
	)
	h := Hash{
		HashName: hashName,
	}

	success, err := h.HsetNx(coon, field, value)
	if err != nil || !success {
		t.Error("hash hsetnx err:", err)
		return
	}

	success, err = h.HsetNx(coon, field, value)
	if err != nil || success {
		t.Error("hash hsetnx err:", err)
		return
	}

	err = Del(coon,h.HashName)
	if err != nil {
		t.Error("hash del err:", err)
		return
	}

	Del(coon,h.HashName)
}

func TestHash_Hincrby(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "age"
		value    = "17"
	)
	h := Hash{
		HashName: hashName,
	}
	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	v,err := h.Hincrby(coon, field, 13)

	if err != nil || v != 30 {
		t.Error("hincrby err:", err)
	}
    t.Log("v:",v)
	Del(coon,h.HashName)
}

func TestHash_HincrbyFloat(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "age"
		value    = 12.8
	)
	h := Hash{
		HashName: hashName,
	}
	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	v, err := h.HincrbyFloat(coon, field, 12.8)
	if err != nil ||  v != 25.6 {
		t.Error("hincrby err:", err)
	}
	t.Log("v:",v)
	Del(coon,h.HashName)
}

func TestHash_HgetAll(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName   = "test"
		filedValue = map[interface{}]interface{}{
			"name": "zhangsan",
			"age":  17,
			"sex":  "man",
		}
	)

	h := Hash{
		HashName: hashName,
	}
	err = h.Hmset(coon, filedValue)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	fv, err := h.HgetAll(coon)
	if err != nil || len(fv) != 3 || fv["name"] != "zhangsan" || fv["age"] != "17" || fv["sex"] != "man" {
		t.Error("hash hget err:", err)
		return
	}

	Del(coon,h.HashName)
}

func TestHash_Hexists(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "name"
		value    = "zhangsan"
	)
	h := Hash{

		HashName: hashName,
	}

	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hset err:", err)
		return
	}

	exist, err := h.Hexists(coon, field)
	if err != nil || !exist {
		t.Error("hexist err:", err, "exist:", exist)
		return
	}

	exist, err = h.Hexists(coon, "aeqiwh")
	if err != nil || exist {
		t.Error("hexist err:", err, "exist:", exist)
		return
	}

	Del(coon,h.HashName)
}

func TestHash_Hlen(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	conn := p.Get()
	defer conn.Close()

	var (
		hashName = "test"
		field = "name"
		value = "zhangsan"
	)
	h := Hash{HashName:hashName}
	err = h.Hset(conn, field, value)
	if err != nil {
		t.Error("hash hset err:", err)
		return
	}

	count, err := h.Hlen(conn)
	if err != nil || count != 1 {
		t.Error("hash hlen err:", err, "count:", count)
		return
	}

	err = Del(conn, h.HashName)
	if err != nil {
		t.Error("hash del err:", err)
		return
	}
}

func TestExistHash(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	conn := p.Get()
	defer conn.Close()

	var (
		hashName = "test"
		field = "name"
		value = "zhangsan"
	)

	h := Hash{HashName:hashName}
	err = h.Hset(conn, field, value)

	err = h.Hset(conn, field, value)
	if err != nil {
		t.Fatal("failed to set hash value,err:", err)
	}
	v, err := ExistKv(conn, hashName)
	if err != nil || !v {
		t.Fatal("failed to get hash value,err:", err)
	}
	err = h.Hdel(conn, field)
	if err != nil {
		t.Fatal("fail to del hash value, err:", err)
	}
}
