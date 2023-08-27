package leetcode

import "testing"

func TestCanCorss(t *testing.T) {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	if !canCross(stones) {
		t.Fatalf("with %v expect true get false", stones)
	}
	stones = []int{0, 1, 2, 3, 4, 8, 9, 11}
	if canCross(stones) {
		t.Fatalf("with %v expect false get true", stones)
	}
}
