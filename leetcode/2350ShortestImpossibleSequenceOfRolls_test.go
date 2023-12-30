package leetcode

import "testing"

func TestShortestSequence(t *testing.T) {
	nums := []int{4, 2, 1, 2, 3, 3, 2, 4, 1}
	k := 4
	if r := shortestSequence(nums, k); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	nums = []int{1, 1, 2, 2}
	k = 2
	if r := shortestSequence(nums, k); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	nums = []int{1, 1, 3, 2, 2, 2, 3, 3}
	k = 4
	if r := shortestSequence(nums, k); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
