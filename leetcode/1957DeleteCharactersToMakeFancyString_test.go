package leetcode

import "testing"

func TestMakeFancyString(t *testing.T) {
	s := "leeetcode"
	exp := "leetcode"
	if r := makeFancyString(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	s = "aaaaa"
	exp = "aa"
	if r := makeFancyString(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	s = "aa"
	if r := makeFancyString(s); r != s {
		t.Fatalf("expect %s get %s", s, r)
	}
}
