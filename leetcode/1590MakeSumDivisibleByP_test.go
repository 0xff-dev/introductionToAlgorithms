package leetcode

import "testing"

func TestMinSubarray(t *testing.T) {
	nums, p, exp := []int{3, 1, 4, 2}, 6, 1
	if r := minSubarray(nums, p); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, p, exp = []int{6, 3, 5, 2}, 9, 2
	if r := minSubarray(nums, p); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, p, exp = []int{1, 2, 3}, 3, 0
	if r := minSubarray(nums, p); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
