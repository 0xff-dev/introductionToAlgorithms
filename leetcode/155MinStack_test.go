package leetcode

import "testing"

func TestMinStack(t *testing.T) {
	c := Constructor4()
	c.Push(-2)
	c.Push(0)
	c.Push(-3)
	if r := c.GetMin(); r != -3 {
		t.Fatalf("expect -3 get %d", r)
	}

	c.Pop()
	if top := c.Top(); top != 0 {
		t.Fatalf("expect 0 get %d", top)
	}

	if r := c.GetMin(); r != -2 {
		t.Fatalf("expect -2 get %d", r)
	}
}
