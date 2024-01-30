package leetcode

import "testing"

func TestNumSubarraysWithSum(t *testing.T) {
	nums := []int{1, 0, 1, 0, 1}
	k := 2
	exp := 4
	if r := numSubarraysWithSum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{0, 0, 0, 0, 0}
	k = 0
	exp = 15
	if r := numSubarraysWithSum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1}
	k = 8
	exp = 20
	if r := numSubarraysWithSum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
