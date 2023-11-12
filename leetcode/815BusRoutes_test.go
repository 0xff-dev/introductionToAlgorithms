package leetcode

import "testing"

func TestNumBusesToDestination(t *testing.T) {
	routes := [][]int{
		{1, 2, 7}, {3, 6, 7},
	}
	source, target := 1, 6
	if r := numBusesToDestination(routes, source, target); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	routes = [][]int{
		{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13},
	}
	source, target = 15, 12
	if r := numBusesToDestination(routes, source, target); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
