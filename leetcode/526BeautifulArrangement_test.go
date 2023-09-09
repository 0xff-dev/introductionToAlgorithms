package leetcode

import "testing"

func TestCountArrangement(t *testing.T) {
	type input struct {
		n, e int
	}
	for _, tc := range []input{
		{1, 1}, {2, 2}, {3, 3}, {5, 10}, {9, 250}, {12, 4010}, {15, 24679},
	} {
		if r := countArrangement(tc.n); r != tc.e {
			t.Fatalf("with input %v expect %d get %d", tc.n, tc.e, r)
		}
	}
}
