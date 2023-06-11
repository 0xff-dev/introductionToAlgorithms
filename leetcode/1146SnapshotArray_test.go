package leetcode

import "testing"

func TestConstructor1146(t *testing.T) {
	c := Constructor1146(3)
	c.Set(0, 5)
	if id := c.Snap(); id != 0 {
		t.Fatalf("expect 0 get %d", id)
	}
	c.Set(0, 6)
	if r := c.Get(0, 0); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
