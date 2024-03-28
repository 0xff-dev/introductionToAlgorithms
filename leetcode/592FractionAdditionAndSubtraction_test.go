package leetcode

import "testing"

func TestFractionAddition(t *testing.T) {
	expression := "-1/2+1/2"
	exp := "0/1"
	if r := fractionAddition(expression); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	expression = "-1/2+1/2+1/3"
	exp = "1/3"
	if r := fractionAddition(expression); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	expression = "1/3-1/2"
	exp = "-1/6"
	if r := fractionAddition(expression); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
