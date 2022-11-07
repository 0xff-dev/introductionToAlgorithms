package leetcode

import "testing"

func TestMaximum69Number(t *testing.T) {
	n := 9669
	if r := maximum69Number(n); r != 9969 {
		t.Fatalf("expect 9969 get %d", r)
	}

	n = 9996
	if r := maximum69Number(n); r != 9999 {
		t.Fatalf("expect 9999 get %d", r)
	}

	n = 9999
	if r := maximum69Number(n); r != 9999 {
		t.Fatalf("expect 9999 get %d", r)
	}
}
