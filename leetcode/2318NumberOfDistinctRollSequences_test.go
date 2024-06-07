package leetcode

import (
	"testing"
)

func TestDistinctSequence(t *testing.T) {
	n := 4
	exp := 184
	if r := distinctSequences(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 2
	exp = 22
	if r := distinctSequences(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 999
	exp = 209513669
	if r := distinctSequences(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
