package leetcode

import "testing"

func TestFind132pattern(t *testing.T) {
	nums := []int{3, 1, 4, 2}
	if !find132pattern(nums) {
		t.Fatalf("expect true get false")
	}
	nums = []int{-1, 3, 2, 0}
	if !find132pattern(nums) {
		t.Fatalf("expect true get false")
	}
	nums = []int{-2, 1, 1}
	if find132pattern(nums) {
		t.Fatalf("expect false get true")
	}
}
