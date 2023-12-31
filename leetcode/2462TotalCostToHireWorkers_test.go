package leetcode

import "testing"

func TestTotalCost(t *testing.T) {
	costs := []int{17, 12, 10, 2, 7, 2, 11, 20, 8}
	k, candidates := 3, 4
	if r := totalCost(costs, k, candidates); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}
	costs = []int{1, 2, 4, 1}
	k, candidates = 3, 3
	if r := totalCost(costs, k, candidates); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
