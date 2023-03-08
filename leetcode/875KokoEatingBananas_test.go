package leetcode

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	piles := []int{3, 6, 7, 11}
	h := 8
	if r := minEatingSpeed(piles, h); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	piles = []int{30, 11, 23, 4, 20}
	h = 5
	if r := minEatingSpeed(piles, h); r != 30 {
		t.Fatalf("expect 30 get %d", r)
	}

	if r := minEatingSpeed(piles, 6); r != 23 {
		t.Fatalf("expect 23 get %d", r)
	}
}
