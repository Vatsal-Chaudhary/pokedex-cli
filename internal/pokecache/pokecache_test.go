package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(5 * time.Minute)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(5 * time.Minute)

	cases := []struct {
		inputKey   string
		inputValue []byte
	}{
		{inputKey: "key1", inputValue: []byte("val1")},
		{inputKey: "key2", inputValue: []byte("val2")},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputValue)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputValue) {
			t.Errorf("%s doesn't match %s", cas.inputKey, cas.inputValue)
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
		t.Errorf("%s should have been reaped", keyOne)
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
		t.Errorf("%s is reaped !!!", keyOne)
	}
}
