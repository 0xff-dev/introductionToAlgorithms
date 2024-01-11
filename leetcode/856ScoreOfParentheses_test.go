package leetcode

import "testing"

func TestScoreOfParenTheses(t *testing.T) {
	s := "()(())"
	if r := scoreOfParentheses(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	s = "()"
	if r := scoreOfParentheses(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	s = "(())"
	if r := scoreOfParentheses(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "()()"
	if r := scoreOfParentheses(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
