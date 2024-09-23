package leetcode

import "testing"

func TestFindNthDigit(t *testing.T) {
	n, exp := 3, 3
	if r := findNthDigit(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, exp = 11, 0
	if r := findNthDigit(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
