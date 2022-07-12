package leetcode

import "testing"

func TestMakesquare(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2}
	if !makesquare(nums) {
		t.Fatalf("expect true get false")
	}

	nums = []int{3, 3, 3, 3, 4}
	if makesquare(nums) {
		t.Fatalf("expect false get true")
	}

	nums = []int{6, 6, 6, 6, 4, 2, 2}
	if makesquare(nums) {
		t.Fatalf("expect false get truej")
	}

	nums = []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}
	if !makesquare(nums) {
		t.Fatalf("expect true get false")
	}
}
