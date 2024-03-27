package leetcode

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	nums := []int{10, 5, 2, 6}
	k := 100
	exp := 8
	if r := numSubarrayProductLessThanK(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1, 2, 3}
	k = 0
	exp = 0
	if r := numSubarrayProductLessThanK(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
