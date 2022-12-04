package leetcode

import "testing"

func TestMinimumAverageDifference(t *testing.T) {
	nums := []int{2, 5, 3, 9, 5, 3}
	if r := minimumAverageDifference(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{0}
	if r := minimumAverageDifference(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
