package leetcode

import "testing"

func TestMinReorder(t *testing.T) {
	n := 6
	connections := [][]int{
		{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5},
	}

	if r := minReorder(n, connections); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	n = 5
	connections = [][]int{
		{1, 0}, {1, 2}, {3, 2}, {3, 4},
	}
	if r := minReorder(n, connections); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	n = 3
	connections = [][]int{
		{1, 0}, {2, 0},
	}
	if r := minReorder(n, connections); r != 0 {
		t.Fatalf("expect 0  get %d", r)
	}
}
