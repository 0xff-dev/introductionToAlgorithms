package leetcode

import "testing"

func TestMaximumSwap(t *testing.T) {
	num, exp := 2736, 7236
	if r := maximumSwap(num); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	num, exp = 9973, 9973
	if r := maximumSwap(num); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	num, exp = 1993, 9913
	if r := maximumSwap(num); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	num, exp = 98368, 98863
	if r := maximumSwap(num); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
