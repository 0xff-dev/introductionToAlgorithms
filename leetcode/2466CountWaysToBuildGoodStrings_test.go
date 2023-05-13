package leetcode

import "testing"

func TestCountGoodStrings(t *testing.T) {
	l, h, z, o := 3, 3, 1, 1
	if r := countGoodStrings(l, h, z, o); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}

	l, h, z, o = 2, 3, 1, 2
	if r := countGoodStrings(l, h, z, o); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
