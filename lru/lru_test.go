package lru

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	lru := New(0, nil)
	lru.Add("key1", "1234")
	if v, ok := lru.Get("key1"); !ok || v.(string) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	lru := New(2, nil)
	lru.Add(k1, v1)
	lru.Add(k2, v2)
	lru.Add(k3, v3)

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value interface{}) {
		keys = append(keys, key)
	}
	lru := New(2, callback)
	lru.Add("key1", "123456")
	lru.Add("k2", "k2")
	lru.Add("k3", "k3")
	lru.Add("k4", "k4")

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
