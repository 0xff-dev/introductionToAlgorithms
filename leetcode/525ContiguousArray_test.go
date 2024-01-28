package leetcode

import "testing"

func TestFindMaxLength(t *testing.T) {
	nums := []int{0, 1}
	if r := findMaxLength(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	nums = []int{0, 1, 0}
	if r := findMaxLength(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
