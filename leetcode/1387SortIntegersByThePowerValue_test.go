package leetcode

import "testing"

func TestGetKth(t *testing.T) {
	lo, hi, k := 12, 15, 2
	if r := getKth(lo, hi, k); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}
}
