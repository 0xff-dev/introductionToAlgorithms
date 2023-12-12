package leetcode

import "testing"

func TestMaximumGap(t *testing.T) {
	nums := []int{3, 6, 9, 1}
	if r := maximumGap(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	nums = []int{100, 3, 2, 1}
	if r := maximumGap(nums); r != 97 {
		t.Fatalf("expect 97 get %d", r)
	}
	nums = []int{3, 1, 4, 9, 10}
	if r := maximumGap(nums); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
