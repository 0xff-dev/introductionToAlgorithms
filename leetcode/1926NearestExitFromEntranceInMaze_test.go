package leetcode

import "testing"

func TestNearestExit(t *testing.T) {
	maze := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}
	entrace := []int{1, 2}
	if r := nearestExit(maze, entrace); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	maze = [][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'},
	}
	entrace = []int{1, 0}
	if r := nearestExit(maze, entrace); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	maze = [][]byte{
		{'.', '+'},
	}

	if r := nearestExit(maze, []int{0, 0}); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
