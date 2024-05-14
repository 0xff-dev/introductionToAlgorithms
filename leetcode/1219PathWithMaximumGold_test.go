package leetcode

import "testing"

func TestGetMaximumGold(t *testing.T) {
	grid := [][]int{
		{0, 6, 0}, {5, 8, 7}, {0, 9, 0},
	}
	exp := 24
	if r := getMaximumGold(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid = [][]int{
		{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20},
	}
	exp = 28
	if r := getMaximumGold(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid = [][]int{
		{1, 0, 7, 0, 0, 0}, {2, 0, 6, 0, 1, 0}, {3, 5, 6, 7, 4, 2}, {4, 3, 1, 0, 2, 0}, {3, 0, 5, 0, 20, 0},
	}
	exp = 60
	if r := getMaximumGold(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
