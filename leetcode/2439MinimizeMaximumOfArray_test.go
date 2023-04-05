package leetcode

import "testing"

func TestMinimizeArrayValue(t *testing.T) {
	nums := []int{3, 7, 1, 6}
	if r := minimizeArrayValue(nums); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums = []int{10, 1}
	if r := minimizeArrayValue(nums); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}
