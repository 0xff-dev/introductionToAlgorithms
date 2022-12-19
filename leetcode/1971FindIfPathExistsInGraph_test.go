package leetcode

import "testing"

func TestValidPath(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}
	sournce, destination := 0, 2
	if !validPath(n, edges, sournce, destination) {
		t.Fatalf("expect true get false")
	}
	if !validPath1(n, edges, sournce, destination) {
		t.Fatalf("expect true get false")
	}

	n = 6
	edges = [][]int{
		{0, 1},
		{0, 2},
		{3, 5},
		{5, 4},
		{4, 3},
	}

	sournce, destination = 0, 5
	if validPath(n, edges, sournce, destination) {
		t.Fatalf("expect false get true")
	}

	if validPath1(n, edges, sournce, destination) {
		t.Fatalf("expect false get true")
	}
}
