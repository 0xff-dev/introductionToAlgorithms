package leetcode

import "testing"

func TestMaxSubarryLength(t *testing.T) {
	nums := []int{1, 2, 3, 1, 2, 3, 1, 2}
	k := 2
	exp := 6
	if r := maxSubarrayLength(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{1, 2, 1, 2, 1, 2, 1, 2}
	k = 1
	exp = 2
	if r := maxSubarrayLength(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{5, 5, 5, 5, 5, 5, 5}
	k = 4
	exp = 4
	if r := maxSubarrayLength(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
