package leetcode

import "testing"

func TestMagicalString(t *testing.T) {
	type input struct {
		n, ans int
	}
	for _, tc := range []input{
		{6, 3}, {7, 4}, {8, 4}, {9, 4}, {10, 5}, {252, 126}, {1314, 659}, {9987, 4990},
	} {
		if r := magicalString(tc.n); r != tc.ans {
			t.Fatalf("with input %d expect %d get %d", tc.n, tc.ans, r)
		}
	}
}
