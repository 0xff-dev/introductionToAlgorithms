package leetcode

import "testing"

func TestLargestSubmatrix(t *testing.T) {
	matrix := [][]int{
		{0, 0, 1}, {1, 1, 1}, {1, 0, 1},
	}
	if r := largestSubmatrix(matrix); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	matrix = [][]int{{1}}
	if r := largestSubmatrix(matrix); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	matrix = [][]int{{0}}
	if r := largestSubmatrix(matrix); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
