package leetcode

import "testing"

func TestMaxSatisfied(t *testing.T) {
	c, g, m := []int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3
	exp := 16
	if r := maxSatisfied(c, g, m); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
