package leetcode

import "testing"

func TestLRUCache(t *testing.T) {
	c := Constructor460(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if r := c.Get(1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	c.h.String()
	c.Put(3, 3)
	c.h.String()
	if r := c.Get(2); r != -1 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := c.Get(3); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	c.Put(4, 4)
	c.h.String()
	if r := c.Get(1); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	if r := c.Get(3); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	if r := c.Get(4); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	t.Log("-------\n")
	c1 := Constructor460(2)
	c1.Put(2, 2)
	c1.Put(1, 1)
	c1.h.String()
	t.Log("\n")
	c1.Put(3, 3)
	c1.h.String()
	if r := c1.Get(1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
