package leetcode

import "testing"

func TestLongestINcreasingPath(t *testing.T) {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	if r := longestIncreasingPath(matrix); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	matrix = [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}

	if r := longestIncreasingPath(matrix); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	matrix = [][]int{{1}}
	if r := longestIncreasingPath(matrix); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
