package leetcode

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	v, exp := []int{8, 1, 5, 2, 6}, 11
	if r := maxScoreSightseeingPair(v); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	v, exp = []int{1, 2}, 2
	if r := maxScoreSightseeingPair(v); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
