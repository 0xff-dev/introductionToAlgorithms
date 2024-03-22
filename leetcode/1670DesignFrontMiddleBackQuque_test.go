package leetcode

import "testing"

func TestConstructor1670(t *testing.T) {
	c := Constructor1670()
	c.PushBack(1)
	c.PushBack(2)
	c.PushMiddle(3)
	c.PushMiddle(4)

	x := c.PopFront()
	if x != 1 {
		t.Fatalf("PopFront expect 1 get %d", x)
	}

	x = c.PopMiddle()
	if x != 3 {
		t.Fatalf("PopMiddle expect 3 get %d", x)
	}

	x = c.PopMiddle()
	if x != 4 {
		t.Fatalf("PopMiddle expect 4 get %d", x)
	}
}
