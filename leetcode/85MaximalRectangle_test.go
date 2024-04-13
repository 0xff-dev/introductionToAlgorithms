package leetcode

import "testing"

func TestMaximalRectangle(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	exp := 6
	if r := maximalRectangle(matrix); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	matrix = [][]byte{
		{'0'},
	}
	exp = 0
	if r := maximalRectangle(matrix); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	matrix = [][]byte{
		{'1'},
	}
	exp = 1
	if r := maximalRectangle(matrix); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
