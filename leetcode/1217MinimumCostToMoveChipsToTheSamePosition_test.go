package leetcode

import "testing"

func TestMinCostToMoveChips(t *testing.T) {
	position := []int{1, 2, 3}
	if minCost := minCostToMoveChips(position); minCost != 1 {
		t.Fatalf("expect 1 get %d", minCost)
	}
	if minCost := minCostToMoveChips2(position); minCost != 1 {
		t.Fatalf("expect 1 get %d", minCost)
	}

	position = []int{2, 2, 2, 3, 3}
	if minCost := minCostToMoveChips(position); minCost != 2 {
		t.Fatalf("expect 2 get %d", minCost)
	}
	if minCost := minCostToMoveChips2(position); minCost != 2 {
		t.Fatalf("expect 2 get %d", minCost)
	}

	position = []int{1, 1000000000}
	if minCost := minCostToMoveChips(position); minCost != 1 {
		t.Fatalf("expect 1 get %d", minCost)
	}
	if minCost := minCostToMoveChips2(position); minCost != 1 {
		t.Fatalf("expect 1 get %d", minCost)
	}
}
