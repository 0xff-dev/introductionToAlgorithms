package leetcode

import "testing"

func TestConstructor307(t *testing.T) {
	c := Constructor307([]int{1, 3, 5})
	if r := c.SumRange(0, 2); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}

	c.Update(1, 2)
	if r := c.SumRange(0, 2); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}

	c = Constructor307([]int{1})
	if r := c.SumRange(0, 1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	c.Update(0, 3)
	if r := c.SumRange(0, 1); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
