package pokecache

import (
	"testing"
	"time"
)

func TestGetSet(t *testing.T) {
	c := NewCache(100 * time.Millisecond)
	c.Set("foo", []byte("bar"))
	val, ok := c.Get("foo")
	if !ok {
		t.Error("expected ok to be true")
	}
	if string(val) != "bar" {
		t.Errorf("expected 'bar', got %q", val)
	}
}

func TestReap(t *testing.T) {
	c := NewCache(100 * time.Millisecond)
	c.Set("foo", []byte("bar"))
	c.reap()
	_, ok := c.Get("foo")
	if !ok {
		t.Error("expected ok to be true")
	}
	time.Sleep(200 * time.Millisecond)
	c.reap()
	_, ok = c.Get("foo")
	if ok {
		t.Error("expected ok to be false")
	}
}

func TestReapLoop(t *testing.T) {
	c := NewCache(100 * time.Millisecond) // reapLoop called internally
	c.Set("foo", []byte("bar"))
	c.Set("baz", []byte("qux"))
	c.Set("quux", []byte("corge"))
	time.Sleep(200 * time.Millisecond)
	_, ok := c.Get("foo")
	if ok {
		t.Error("expected ok to be false")
	}
	_, ok = c.Get("baz")
	if ok {
		t.Error("expected ok to be false")
	}
	_, ok = c.Get("quux")
	if ok {
		t.Error("expected ok to be false")
	}
}
