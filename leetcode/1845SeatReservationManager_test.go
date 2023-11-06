package leetcode

import "testing"

func TestConstructor1845(t *testing.T) {
	c := Constructor1845(5)
	a := c.Reserve()
	b := c.Reserve()
	if a != 1 || b != 2 {
		t.Fatalf("fail")
	}
	c.Unreserve(2)
	d := c.Reserve()
	e := c.Reserve()
	f := c.Reserve()
	g := c.Reserve()
	if d != 2 || e != 3 || f != 4 || g != 5 {
		t.Fatalf("fail")
	}
	c.Unreserve(5)
}
