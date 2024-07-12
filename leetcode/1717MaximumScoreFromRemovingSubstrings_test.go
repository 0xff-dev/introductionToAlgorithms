package leetcode

import "testing"

func TestMaximumGain(t *testing.T) {
	s := "cdbcbbaaabab"
	x, y := 4, 5
	exp := 19
	if r := maximumGain(s, x, y); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
