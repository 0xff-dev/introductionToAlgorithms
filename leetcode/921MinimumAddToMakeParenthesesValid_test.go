package leetcode

import "testing"

func TestMinAddmakeValid(t *testing.T) {
	s := "())"
	if r := minAddToMakeValid(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	s = "((("
	if r := minAddToMakeValid(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
