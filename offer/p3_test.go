package offer

import "testing"

func TestFind(t *testing.T) {
	input := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	if Find(input, len(input), len(input[0]), 7) != true {
		t.Fatal("except true")
	}
}
