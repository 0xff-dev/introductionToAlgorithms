package leetcode

import "testing"

func TestFindClosestNumber(t *testing.T) {
	nums := []int{2, -1, 1}
	if r := findClosestNumber(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
