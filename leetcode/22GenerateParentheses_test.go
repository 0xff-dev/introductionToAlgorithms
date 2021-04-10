package leetcode

import "testing"

func TestGenerateParentheses(t *testing.T) {
	for n := 1; n <= 8; n++ {
		r := generateParenthesis(n)
		t.Logf("%v\n", r)
	}
}
