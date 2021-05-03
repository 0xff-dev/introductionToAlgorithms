package leetcode

import "testing"

func TestMoveZeroes(t *testing.T) {
	for _, testCase := range [][]int{
		{0, 1, 0, 3, 13},
		{1, 0, 0, 0, 3, 0, 0, 2, 0, 4},
		{0, 0, 0},
		{0},
		{1},
		{1, 3, 5},
		{0, 0, 0, 1, 0, 2, 0, 3, 0},
		{1, 1, 0, 1},
	} {
		t.Logf("source: %v", testCase)
		moveZeroes(testCase)
		t.Logf("result: %v", testCase)
	}
}
