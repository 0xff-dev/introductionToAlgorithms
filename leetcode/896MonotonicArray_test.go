package leetcode

import "testing"

func TestIsMonotonic(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	if !isMonotonic(nums) {
		t.Fatalf("expect true get false")
	}
	if !isMonotonic2(nums) {
		t.Fatalf("expect true get false")
	}
	nums = []int{6, 5, 4, 4}
	if !isMonotonic(nums) {
		t.Fatalf("expect true get false")
	}
	if !isMonotonic2(nums) {
		t.Fatalf("expect true get false")
	}

	nums = []int{1, 3, 2}
	if isMonotonic(nums) {
		t.Fatalf("expect false get true")
	}
	if isMonotonic2(nums) {
		t.Fatalf("expect false get true")
	}
}
