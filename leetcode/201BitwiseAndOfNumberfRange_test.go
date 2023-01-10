package leetcode

import "testing"

func TestRangeBitwistAnd(t *testing.T) {
	left, right := 5, 7
	if r := rangeBitwiseAnd(left, right); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	left, right = 0, 0
	if r := rangeBitwiseAnd(left, right); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
