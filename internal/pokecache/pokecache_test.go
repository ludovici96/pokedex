
package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cache := NewCache(interval)

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache.Add(c.key, c.val)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Fatal("expected to find key")
			}

			if string(val) != string(c.val) {
				t.Errorf("expected %s, got %s", c.val, val)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 50 * time.Millisecond
	const waitTime = baseTime + 50*time.Millisecond

	cache := NewCache(baseTime)

	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Fatal("expected to find key before reaping")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Fatal("expected key to be removed")
	}
}
