package leetcode

import "testing"

func TestMonotoneIncreasingDigits(t *testing.T) {
	n := 10
	exp := 9
	if r := monotoneIncreasingDigits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 1234
	exp = 1234
	if r := monotoneIncreasingDigits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 332
	exp = 299
	if r := monotoneIncreasingDigits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
