package leetcode

import "testing"

func TestMaxCoins(t *testing.T) {
	piles := []int{2, 4, 1, 2, 7, 8}
	if r := maxCoins1561(piles); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
	piles = []int{2, 4, 5}
	if r := maxCoins1561(piles); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	piles = []int{9, 8, 7, 6, 5, 1, 2, 3, 4}
	if r := maxCoins1561(piles); r != 18 {
		t.Fatalf("expect 18 get %d", r)
	}
}
