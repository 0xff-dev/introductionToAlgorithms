package leetcode

import "testing"

func TestOriginalDigits(t *testing.T) {
	s := "owoztneoer"
	exp := "012"
	if r := originalDigits(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s = "fviefuro"
	exp = "45"
	if r := originalDigits(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
