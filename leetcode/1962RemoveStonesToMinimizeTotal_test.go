package leetcode

import "testing"

func TestMinStoneSum(t *testing.T) {
	piles, k := []int{5, 4, 9}, 2
	if r := minStoneSum(piles, k); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	piles, k = []int{4, 3, 6, 7}, 3
	if r := minStoneSum(piles, k); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
}
