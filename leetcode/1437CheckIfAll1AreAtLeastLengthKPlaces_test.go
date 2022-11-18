package leetcode

import "testing"

func TestKLengthPart(t *testing.T) {
	nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	k := 2
	if !kLengthApart(nums, k) {
		t.Fatalf("expect true get false")
	}

	nums = []int{1, 0, 0, 1, 0, 1}
	if kLengthApart(nums, k) {
		t.Fatalf("expect false get true")
	}
}
