package leetcode

import "testing"

func TestMinimumCost(t *testing.T) {
	/*
		source, target, o, c, cost, exp := "abcd", "acbe", []byte("abcced"), []byte("bcbebe"), []int{2, 5, 5, 1, 2, 20}, int64(28)
		if r := minimumCost(source, target, o, c, cost); r != exp {
			t.Fatalf("expect %d get %d", exp, r)
		}
	*/

	source, target, o, c, cost, exp := "aaaa", "bbbb", []byte("ac"), []byte("cb"), []int{1, 2}, int64(12)
	if r := minimumCost(source, target, o, c, cost); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	source, target, o, c, cost, exp = "abcd", "abce", []byte("a"), []byte("e"), []int{10000}, -1

	if r := minimumCost(source, target, o, c, cost); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
