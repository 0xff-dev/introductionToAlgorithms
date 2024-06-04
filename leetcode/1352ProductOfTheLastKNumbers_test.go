package leetcode

import "testing"

func TestConstructor1352(t *testing.T) {
	c := Constructor1352()
	c.Add(3)
	c.Add(0)
	c.Add(2)
	c.Add(5)
	c.Add(4)
	if r := c.GetProduct(2); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}
	if r := c.GetProduct(3); r != 40 {
		t.Fatalf("expect 40 get %d", r)
	}
	if r := c.GetProduct(4); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	c.Add(8)
	if r := c.GetProduct(2); r != 32 {
		t.Fatalf("expect 32 get %d", r)
	}
}
