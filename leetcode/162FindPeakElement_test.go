package leetcode

import "testing"

func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	if r := findPeakelement(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	if r := findPeakelement(nums); !(r == 1 || r == 5) {
		t.Fatalf("expect 1 or 5 get %d", r)
	}
}
