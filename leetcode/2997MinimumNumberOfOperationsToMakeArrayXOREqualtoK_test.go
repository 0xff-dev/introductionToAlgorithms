package leetcode

import "testing"

func TestMinOperations2997(t *testing.T) {
	nums := []int{2, 1, 3, 4}
	k := 1
	exp := 2
	if r := minOperations2997(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{2, 0, 2, 0}
	k = 0
	exp = 0
	if r := minOperations2997(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
