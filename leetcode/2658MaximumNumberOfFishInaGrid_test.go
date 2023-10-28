package leetcode

import "testing"

func TestFindMaxFish(t *testing.T) {
	grid := [][]int{
		{0, 2, 1, 0},
		{4, 0, 0, 3},
		{1, 0, 0, 4},
		{0, 3, 2, 0},
	}
	if r := findMaxFish(grid); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	grid = [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
	}
	if r := findMaxFish(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
