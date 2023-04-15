package leetcode

import "testing"

func TestMaxValueOfConis(t *testing.T) {
	piles := [][]int{
		{1, 100, 3}, {7, 8, 9},
	}
	if r := maxValueOfCoins(piles, 2); r != 101 {
		t.Fatalf("expect 101 get %d", r)
	}
}
