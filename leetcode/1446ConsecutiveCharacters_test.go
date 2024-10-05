package leetcode

import "testing"

func TestMaxPower(t *testing.T) {
	s, exp := "leetcode", 2
	if r := maxPower(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "abbcccddddeeeeedcba", 5
	if r := maxPower(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
