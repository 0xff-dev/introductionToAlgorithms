package leetcode

import "testing"


func TestCheck(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	if !check(nums) {
		t.Fatalf("expect true get false")
	}

	nums = []int{2, 1, 3, 4}
	if check(nums) {
		t.Fatalf("expect false get true")
	}

	nums = []int{1, 3, 2}
	if check(nums) {
		t.Fatalf("expect false get true")
	}
}

