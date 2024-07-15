package leetcode

import "testing"

func TestParseBoolExpr(t *testing.T) {
	e := "&(|(f))"
	exp := false
	if r := parseBoolExpr(e); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	e = "|(f,f,f,t)"
	exp = true
	if r := parseBoolExpr(e); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	e = "!(&(f,t))"
	exp = true
	if r := parseBoolExpr(e); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	e = "|(&(t,f,t),!(t))"
	exp = false
	if r := parseBoolExpr(e); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
