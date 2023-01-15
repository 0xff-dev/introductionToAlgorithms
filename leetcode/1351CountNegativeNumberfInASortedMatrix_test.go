package leetcode

import "testing"

func TestCountNetagive(t *testing.T) {
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	if r := countNegatives(grid); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}
