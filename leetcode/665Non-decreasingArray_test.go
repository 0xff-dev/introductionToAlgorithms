package leetcode

import "testing"

func TestCheckPossibility(t *testing.T) {
	nums := []int{4, 2, 3}
	if !checkPossibility(nums) {
		t.Fatalf("expect true get false")
	}

	nums = []int{4, 2, 1}
	if checkPossibility(nums) {
		t.Fatalf("expect false get true")
	}

	nums = []int{3, 4, 2, 3}
	if checkPossibility(nums) {
		t.Fatalf("expect false get true")
	}

	nums = []int{-1, 4, 2, 3}
	if !checkPossibility(nums) {
		t.Fatalf("expect true get false")
	}
}
