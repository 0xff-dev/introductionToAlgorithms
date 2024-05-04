package leetcode

import "testing"

func TestMaxEnvelopes(t *testing.T) {
	envelopes := [][]int{
		{5, 4}, {6, 4}, {6, 7}, {2, 3},
	}
	exp := 3
	if r := maxEnvelopes(envelopes); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	envelopes = [][]int{
		{1, 1}, {1, 1}, {1, 1},
	}
	exp = 1
	if r := maxEnvelopes(envelopes); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
