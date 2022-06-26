package leetcode

import "testing"

func TestJudgeCircle(t *testing.T) {
	moves := "UD"
	if !judgeCircle(moves) {
		t.Fatal("expect true get false")
	}

	moves = "LL"
	if judgeCircle(moves) {
		t.Fatal("expect false get true")
	}
}