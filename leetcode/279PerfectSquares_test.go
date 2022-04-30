package leetcode

import "testing"

func TestNumSquares(t *testing.T) {
	n := 12
	if r := numSquares(n); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	if r := numSquares(13); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	if r := numSquares(268); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
