package leetcode

import "testing"

type input901 struct {
	n, exp int
}

func TestConstructor901(t *testing.T) {
	c := Constructor901()

	for _, tc := range []input901{
		{100, 1},
		{80, 1},
		{60, 1},
		{70, 2},
		{60, 1},
		{75, 4},
		{85, 6},
	} {
		if r := c.Next(tc.n); r != tc.exp {
			t.Fatalf("with n %d expect %d get %d", tc.n, tc.exp, r)
		}
	}

}
