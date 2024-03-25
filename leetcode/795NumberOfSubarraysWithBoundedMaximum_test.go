package leetcode

import "testing"

func TestNumSubarrayBoundedMax(t *testing.T) {
	nums := []int{2, 1, 4, 3}
	l, r := 2, 3
	exp := 3
	if ans := numSubarrayBoundedMax(nums, l, r); ans != exp {
		t.Fatalf("expect %d get %d", exp, ans)
	}
	nums = []int{2, 9, 2, 5, 6}
	l, r = 2, 8
	exp = 7
	if ans := numSubarrayBoundedMax(nums, l, r); ans != exp {
		t.Fatalf("expect %d get %d", exp, ans)
	}
}
