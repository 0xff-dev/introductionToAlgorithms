package leetcode

import "testing"

func TestCountOrders(t *testing.T) {
	type input struct {
		n   int
		exp int
	}
	for _, tc := range []input{
		{1, 1}, {2, 6}, {3, 90}, {128, 761756041}, {345, 432656738}, {491, 452524619}, {500, 764678010},
	} {
		if r := countOrders(tc.n); r != tc.exp {
			t.Fatalf("with %d expect %d get %d", tc.n, tc.exp, r)
		}
	}
}
