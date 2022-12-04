package leetcode

import "testing"

func TestNumSpecial(t *testing.T) {
	mat := [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}

	if r := numSpecial(mat); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	mat = [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	if r := numSpecial(mat); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
