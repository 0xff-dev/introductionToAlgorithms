package leetcode

import "testing"

func TestKnightProbability(t *testing.T) {
	n, k, r, c := 3, 2, 0, 0
	if r := knightProbability(n, k, r, c); r != 0.0625 {
		t.Fatalf("expect 0.0625 get %f", r)
	}
}
