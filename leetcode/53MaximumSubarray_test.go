package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if r := maxSubArray(input); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	input = []int{1}
	if r := maxSubArray(input); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	input = []int{5, 4, -1, 7, 8}
	if r := maxSubArray(input); r != 23 {
		t.Fatalf("expect 23 get %d", r)
	}
}
