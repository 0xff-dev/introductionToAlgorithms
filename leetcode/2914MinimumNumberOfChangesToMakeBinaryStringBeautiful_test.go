package leetcode

import "testing"

func TestMinChanges(t *testing.T) {
	s, exp := "1001", 2
	if r := minChanges(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "10", 1
	if r := minChanges(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "0000", 0
	if r := minChanges(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
