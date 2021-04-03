package leetcode

import "testing"

func TestMnCostClimbingStairs(t *testing.T) {
	input := []int{10, 15, 20}
	if r := minCostClimbingStairs(input); r != 15 {
		t.Fatalf("expect 15 get %d", r)
	}

	input = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	if r := minCostClimbingStairs(input); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
