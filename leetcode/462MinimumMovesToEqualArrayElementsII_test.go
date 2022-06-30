package leetcode

import "testing"

func TestMinMovex2(t *testing.T) {
	nums := []int{1, 2, 3}
	if r := minMoves2(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{1, 10, 2, 9}
	if r := minMoves2(nums); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}
}
