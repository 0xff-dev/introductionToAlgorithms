package leetcode

import "testing"

func TestMinSwap1963(t *testing.T) {
	s, exp := "][][", 1
	if r := minSwaps1963(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "]]][[[", 2
	if r := minSwaps1963(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "[]", 0
	if r := minSwaps1963(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
