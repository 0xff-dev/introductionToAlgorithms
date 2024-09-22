package leetcode

import "testing"

func TestFindKthNumber(t *testing.T) {
	n, k, exp := 13, 2, 10
	if r := findKthNumber(n, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, k, exp = 1, 1, 1
	if r := findKthNumber(n, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	n, k, exp = 156, 132, 77
	if r := findKthNumber(n, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, k, exp = 100, 10, 17
	if r := findKthNumber(n, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
