package leetcode

import "testing"

func MinimumSwap(t *testing.T) {
	s1, s2 := "xx", "yy"
	if r := minimumSwap(s1, s2); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	s1, s2 = "xy", "yx"
	if r := minimumSwap(s1, s2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s1, s2 = "xx", "xy"
	if r := minimumSwap(s1, s2); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
