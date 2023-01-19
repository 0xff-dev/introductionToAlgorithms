package leetcode

import "testing"

func TestSubarrayDivByK(t *testing.T) {
	nums, k := []int{4, 5, 0, -2, -3, 1}, 5
	if r := subarraysDivByK(nums, k); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	nums, k = []int{5}, 9
	if r := subarraysDivByK(nums, k); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
