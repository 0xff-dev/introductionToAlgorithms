package leetcode

import "testing"

func TestCountSubarrays2962(t *testing.T) {
	nums := []int{1, 3, 2, 3, 3}
	k := 2
	exp := int64(6)
	if r := countSubarrays2962(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{1, 4, 2, 1}
	k = 3
	exp = 0
	if r := countSubarrays2962(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
