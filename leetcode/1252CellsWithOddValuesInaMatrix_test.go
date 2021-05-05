package leetcode

import "testing"

func TestOddCells(t *testing.T) {
	m, n := 2, 3
	indices := [][]int{
		{0, 1},
		{1, 1},
	}

	if r := oddCells(m, n, indices); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	m, n = 2, 2
	indices = [][]int{
		{1, 1},
		{0, 0},
	}

	if r := oddCells(m, n, indices); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	if r := oddCells(m, n, [][]int{}); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
