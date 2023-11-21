package leetcode

import "testing"

func TestCountNicePairs(t *testing.T) {
	nums := []int{42, 11, 1, 97}
	if r := countNicePairs(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	nums = []int{13, 10, 35, 24, 76}
	if r := countNicePairs(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
