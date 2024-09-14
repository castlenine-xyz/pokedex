package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "Key1",
			inputVal: []byte("Val1"),
		},
		{
			inputKey: "Key2",
			inputVal: []byte("Val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("Val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s doesn't match %s", string(actual), string(cas.inputVal))
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"

	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should've been reaped", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"

	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval / 2)
	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s shouldn't be reaped", keyOne)
	}
}
