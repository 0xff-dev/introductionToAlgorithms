package leetcode

import "testing"

type input304 struct {
	row1, col1, row2, col2 int
	ans                    int
}

func TestConstructor304(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	o := Constructor304(matrix)
	for _, tc := range []input304{
		{2, 1, 4, 3, 8},
		{1, 1, 2, 2, 11},
		{1, 2, 2, 4, 12},
		{0, 0, 0, 0, 3},
		{4, 4, 4, 4, 5},
	} {
		if r := o.SumRegion(tc.row1, tc.col1, tc.row2, tc.col2); r != tc.ans {
			t.Fatalf("expect %d get %d", tc.ans, r)
		}
	}

	matrix = [][]int{
		{-4, -5},
	}
	o = Constructor304(matrix)
	for _, tc := range []input304{
		{0, 0, 0, 0, -4},
		{0, 0, 0, 1, -9},
		{0, 1, 0, 1, -5},
	} {
		if r := o.SumRegion(tc.row1, tc.col1, tc.row2, tc.col2); r != tc.ans {
			t.Fatalf("expect %d get %d", tc.ans, r)
		}
	}
}
