package leetcode

import "testing"

func TestStoneGame(t *testing.T) {
	piles := []int{5, 3, 4, 5}
	if r := stoneGame(piles); !r {
		t.Fatalf("expect true get false")
	}

}
