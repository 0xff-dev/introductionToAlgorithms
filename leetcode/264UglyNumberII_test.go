package leetcode

import "testing"

func TestUglyNumber(t *testing.T) {
	if r := nthUglyNumber(1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	if r := nthUglyNumber(1690); r != 2123366400 {
		t.Fatalf("expect 2123366400 get %d", r)
	}
}
