package leetcode

import "testing"

func TestReverseParentheses(t *testing.T) {
	s := "(u(love)i)"
	exp := "iloveu"
	if r := reverseParentheses(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s = "(abcd)"
	exp = "dcba"
	if r := reverseParentheses(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s = "(ed(et(oc))el)"
	exp = "leetcode"
	if r := reverseParentheses(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
