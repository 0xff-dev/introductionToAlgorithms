package leetcode

import "testing"

func TestRemoveOccurences(t *testing.T) {
	s := "eemckxmckx"
	part := "emckx"
	exp := ""
	if r := removeOccurrences(s, part); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
