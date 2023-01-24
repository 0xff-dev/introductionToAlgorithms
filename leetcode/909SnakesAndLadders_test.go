package leetcode

import "testing"

func TestSnakesAndLadders(t *testing.T) {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}
	if r := snakesAndLadders(board); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	board = [][]int{
		{-1, -1, -1},
		{-1, 9, 8},
		{-1, 8, 9},
	}
	if r := snakesAndLadders(board); r != 1 {
		t.Fatalf("expect 3 get %d", r)
	}
	board = [][]int{
		{1, 1, -1},
		{1, 1, 1},
		{-1, 1, 1},
	}
	if r := snakesAndLadders(board); r != -1 {
		t.Fatalf("expect 2 get %d", r)
	}
}
