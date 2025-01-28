package leetcode

import "testing"

func TestFindIntegers(t *testing.T) {
	n, exp := 1000000000, 2178309
	if r := findIntegers(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
