package leetcode

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {
	s := "lee(t(c)o)de)"
	exp := "lee(t(c)o)de"
	if r := minRemoveToMakeValid(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s = "a)b(c)d"
	exp = "ab(c)d"
	if r := minRemoveToMakeValid(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s = ")))((("
	exp = ""
	if r := minRemoveToMakeValid(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
