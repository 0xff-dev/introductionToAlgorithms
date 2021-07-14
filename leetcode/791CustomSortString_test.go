package leetcode

import "testing"

func TestCustomSortString(t *testing.T) {
	order, str := "cba", "abcd"
	r := customSortString(order, str)
	t.Logf("res: %s", r)
	r1 := customSortString1(order, str)
	t.Logf("res1: %s", r1)
}
