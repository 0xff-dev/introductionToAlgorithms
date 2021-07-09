package leetcode

import "testing"

func TestSearchMatrixII(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if r := searchMatrixII(matrix, 5); !r {
		t.Fatalf("expect true get false")
	}

	matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if r := searchMatrixII(matrix, 20); r {
		t.Fatalf("expect false get true")
	}

	matrix = [][]int{
		{1},
	}
	if r := searchMatrixII(matrix, 1); !r {
		t.Fatalf("expect true get false")
	}

	if r := searchMatrixII(matrix, 2); r {
		t.Fatalf("expect false get true")
	}

	matrix = [][]int{
		{1, 2, 3},
	}

	if r := searchMatrixII(matrix, 2); !r {
		t.Fatalf("expect true get false")
	}

	if r := searchMatrixII(matrix, 1); !r {
		t.Fatalf("expect true get false")
	}

	matrix = [][]int{
		{1},
		{2},
		{3},
	}

	if r := searchMatrixII(matrix, 1); !r {
		t.Fatalf("expect true get flase")
	}

	if r := searchMatrixII(matrix, 4); r {
		t.Fatalf("expect fasle get true")
	}

	if r := searchMatrixII(matrix, 2); !r {
		t.Fatalf("expect true get false")
	}
}
