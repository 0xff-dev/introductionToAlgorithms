package leetcode

import "testing"

func TestMaxScore1799(t *testing.T) {
	nums := []int{1, 2}
	if r := maxScore1799(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	nums = []int{3, 4, 6, 8}
	if r := maxScore1799(nums); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}
	nums = []int{415, 230, 471, 705, 902, 87}
	if r := maxScore1799(nums); r != 23 {
		t.Fatalf("expect 23 get %d", r)
	}
}
