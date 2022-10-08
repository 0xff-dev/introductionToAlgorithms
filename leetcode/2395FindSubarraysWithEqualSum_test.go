package leetcode

import "testing"

func TestFindSubarrays(t *testing.T) {
	nums := []int{4, 2, 4}
	if !findSubarrays(nums) {
		t.Fatalf("expect true get false")
	}

	nums = []int{1, 2, 3, 4, 5}
	if findSubarrays(nums) {
		t.Fatalf("expect false get true")
	}

	nums = []int{0, 0, 0}
	if !findSubarrays(nums) {
		t.Fatalf("expect true get false")
	}
}
