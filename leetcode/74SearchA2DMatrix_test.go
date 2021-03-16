package leetcode

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	input := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	if !searchMatrix(input, 3) {
		t.Fatalf("expect true get false")
	}

	if searchMatrix(input, 13) {
		t.Fatalf("expect false get ture")
	}

	input = [][]int{
		{1},
	}
	if !searchMatrix(input, 1) {
		t.Fatalf("expect treu get false")
	}
	if searchMatrix(input, 2) {
		t.Fatalf("expect false get true")
	}
}
