package leetcode

import "testing"

func TestFindPairs(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5}
	k := 2
	exp := 2
	if r := findPairs(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1, 2, 3, 4, 5}
	k = 1
	exp = 4
	if r := findPairs(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1, 3, 1, 5, 4}
	k = 0
	exp = 1
	if r := findPairs(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
