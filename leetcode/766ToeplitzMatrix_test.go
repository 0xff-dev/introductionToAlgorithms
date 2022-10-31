package leetcode

import "testing"

func TestIsToeplitzMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}

	if !isToeplitzMatrix(matrix) {
		t.Fatalf("expect true get false")
	}

	matrix = [][]int{
		{1, 2},
		{2, 2},
	}

	if isToeplitzMatrix(matrix) {
		t.Fatalf("expect false get true")
	}

	matrix = [][]int{
		{1},
	}

	if !isToeplitzMatrix(matrix) {
		t.Fatalf("expect true get false")
	}
}
