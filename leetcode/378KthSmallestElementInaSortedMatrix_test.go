package leetcode

import "testing"

func TestKthSmallest(t *testing.T) {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	if r := kthSmallest(matrix, 8); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}

	matrix = [][]int{
		{-5},
	}

	if r := kthSmallest(matrix, 1); r != -5 {
		t.Fatalf("expect -5 get %d", r)
	}
}
