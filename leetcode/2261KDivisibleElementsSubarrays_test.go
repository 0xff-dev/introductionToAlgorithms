package leetcode

import "testing"

func TestCountDistinct(t *testing.T) {
	nums := []int{2, 3, 3, 2, 2}
	k, p := 2, 2
	if r := countDistinct(nums, k, p); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}
	nums = []int{1, 2, 3, 4}
	k, p = 4, 1
	if r := countDistinct(nums, k, p); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	nums = []int{6, 20, 5, 18}
	k, p = 3, 14
	if r := countDistinct(nums, k, p); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}
