package leetcode

import "testing"

func TestIntegerBreak(t *testing.T) {
	n := 2
	if r := integerBreak(n); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := integerBreak343(n); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	n = 10
	if r := integerBreak(n); r != 36 {
		t.Fatalf("expect 36 get %d", r)
	}
	if r := integerBreak343(n); r != 36 {
		t.Fatalf("expect 36 get %d", r)
	}
}
