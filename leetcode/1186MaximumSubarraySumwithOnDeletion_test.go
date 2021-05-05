package leetcode

import "testing"

func TestMaximumSum(t *testing.T) {
	input := []int{1, -2, 0, 3}
	if r := maximumSum(input); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	input = []int{1, -2, -2, 3}
	if r := maximumSum(input); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	input = []int{-1, -1, -1, -1}
	if r := maximumSum(input); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	input = []int{2, 1, -2, -5, -2}
	if r := maximumSum(input); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
