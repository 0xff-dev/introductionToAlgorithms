package leetcode

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	n := 5
	if !judgeSquareSum(n) {
		t.Fatalf("expect true get false")
	}
	n = 3
	if judgeSquareSum(n) {
		t.Fatalf("expect false get true")
	}
}
