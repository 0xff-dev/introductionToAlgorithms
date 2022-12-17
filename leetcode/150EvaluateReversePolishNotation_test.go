package leetcode

import "testing"

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	if r := evalRPN(tokens); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}

	tokens = []string{"4", "13", "5", "/", "+"}
	if r := evalRPN(tokens); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	if r := evalRPN(tokens); r != 22 {
		t.Fatalf("expect 22 get %d", r)
	}
}
