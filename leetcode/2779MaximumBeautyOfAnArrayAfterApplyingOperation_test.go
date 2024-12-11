package leetcode

import "testing"

func TestMaximumBeauty2779(t *testing.T) {
	nums, k, exp := []int{4, 6, 1, 2}, 2, 3
	if r := maximumBeauty2779(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{1, 1, 1, 1}, 10, 4
	if r := maximumBeauty2779(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
