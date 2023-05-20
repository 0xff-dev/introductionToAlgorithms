package leetcode

import "testing"

func TestSubrectangleQueries(t *testing.T) {
	c := Constructor1476([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	if r := c.GetValue(0, 2); r != 1 {
		t.Fatalf("expect 2 get %d", r)
	}

	c.UpdateSubrectangle(0, 0, 3, 2, 5)
	if r := c.GetValue(0, 2); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	if r := c.GetValue(3, 1); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	c.UpdateSubrectangle(3, 0, 3, 2, 10)
	if r := c.GetValue(3, 1); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	if r := c.GetValue(0, 2); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
