package leetcode

import "testing"

func TestMinSumPath(t *testing.T) {
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	if r := minPathSum(input); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	input = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	if r := minPathSum(input); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	input = [][]int{
		{1},
	}
	if r := minPathSum(input); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	input = [][]int{
		{1, 2},
	}
	if r := minPathSum(input); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	input = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
		{1, 1, 1},
	}
	if r := minPathSum(input); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}
