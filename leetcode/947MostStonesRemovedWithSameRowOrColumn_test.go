package leetcode

import "testing"

func TestRemoveStones(t *testing.T) {
	stones := [][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},
		{2, 2},
	}

	if r := removeStones(stones); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	stones = [][]int{
		{0, 0},
		{0, 2},
		{1, 1},
		{2, 0},
		{2, 2},
	}
	if r := removeStones(stones); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	stones = [][]int{
		{0, 0},
	}

	if r := removeStones(stones); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	stones = [][]int{
		{3, 2},
		{3, 1},
		{4, 4},
		{1, 1},
		{0, 2},
		{4, 0},
	}

	if r := removeStones(stones); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
