package leetcode

import "testing"

func TestMinimumFuelCost(t *testing.T) {
	roads, seats := [][]int{{0, 1}, {0, 2}, {0, 3}}, 5
	if r := minimumFuelCost(roads, seats); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	roads, seats = [][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2
	if r := minimumFuelCost(roads, seats); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	roads, seats = [][]int{}, 1
	if r := minimumFuelCost(roads, seats); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
