package challengeprogrammingdatastructure

import "testing"

func TestLargestSquare(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
	}
	if r := LargestSquare(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	grid = [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	}
	if r := LargestSquare(grid); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}
}
