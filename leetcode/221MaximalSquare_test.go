package leetcode

import "testing"

func TestMaximalSquare(t *testing.T) {
	m := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	if r := maximalSquare(m); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	m = [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	if r := maximalSquare(m); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
