package leetcode

import "testing"

func TestRotateDigits(t *testing.T) {
	n := 10
	exp := 4
	if r := rotatedDigits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	n = 1
	exp = 0
	if r := rotatedDigits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 2
	exp = 1
	if r := rotatedDigits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
