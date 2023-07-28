package leetcode

import "testing"

func TestPredictTheWinner(t *testing.T) {
	if PredictTheWinner([]int{1, 5, 2}) {
		t.Fatalf("expect false get true")
	}
}
