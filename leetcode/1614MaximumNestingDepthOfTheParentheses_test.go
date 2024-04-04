package leetcode

import "testing"

func TestMaxDepth1614(t *testing.T) {
	s := "(1+(2*3)+((8)/4))+1"
	exp := 3
	if r := maxDepth1614(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	s = "(1)+((2))+(((3)))"
	if r := maxDepth1614(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
