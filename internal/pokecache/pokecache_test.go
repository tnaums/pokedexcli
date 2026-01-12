package pokecache

import (
    "testing"
    "time"
)

func TestAddCache(t *testing.T) {
    interval := 5 * time.Second

    key := "https://example.com"
    val := []byte("testdata")

    cache := NewCache(interval)

    cache.Add(key, val)

    got, ok := cache.Get(key)
    if !ok {
        t.Errorf("expected key %q to be found", key)
        return
    }

    if string(got) != string(val) {
        t.Errorf("expected value %q, got %q", string(val), string(got))
    }
}
