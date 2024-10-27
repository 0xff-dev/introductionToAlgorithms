package leetcode

import "testing"

func TestClumsy(t *testing.T) {
	n, exp := 4, 7
	if r := clumsy(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, exp = 10, 12
	if r := clumsy(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
