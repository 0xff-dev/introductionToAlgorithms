package leetcode

import "testing"

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	if r := pivotIndex(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{1, 2, 3}
	if r := pivotIndex(nums); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	nums = []int{2, 1, -1}
	if r := pivotIndex(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
