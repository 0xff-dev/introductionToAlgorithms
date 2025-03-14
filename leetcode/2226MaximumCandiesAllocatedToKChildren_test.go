package leetcode

import "testing"

func TestMaximumCandies(t *testing.T) {
	candies, k, exp := []int{5, 8, 6}, int64(3), 5
	if r := maximumCandies(candies, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	candies, k, exp = []int{2, 5}, 11, 0
	if r := maximumCandies(candies, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
