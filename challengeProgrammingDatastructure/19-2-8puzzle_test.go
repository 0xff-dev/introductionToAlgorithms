package challengeprogrammingdatastructure

import "testing"

func TestPuzzle(t *testing.T) {
	grid := [][]int{
		{1, 3, 0},
		{4, 2, 5},
		{7, 8, 6},
	}
	if r := Puzzle(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
