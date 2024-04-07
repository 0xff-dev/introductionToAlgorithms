package leetcode

import "testing"

func TestCheckValidString(t *testing.T) {
	s := "("
	exp := false
	if r := checkValidString(s); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	s = "************************************************************"
	exp = true
	if r := checkValidString(s); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
