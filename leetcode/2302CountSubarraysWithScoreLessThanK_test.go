package leetcode

import "testing"

func TestCountSubarrays2302(t *testing.T) {
	nums := []int{2, 1, 4, 3, 5}
	k := int64(10)
	exp := int64(6)
	if r := countSubarrays2302(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1, 1, 1}
	k = 5
	exp = 5
	if r := countSubarrays2302(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
