package leetcode

import "testing"

type TrailingZeroCase struct {
	N   int
	Ans int
}

func TestTrailingZeroes(t *testing.T) {
	for _, tc := range []TrailingZeroCase{
		{
			N:   50,
			Ans: 12,
		},
		{
			N:   3,
			Ans: 0,
		},
		{
			N:   5,
			Ans: 1,
		},
		{
			N:   937,
			Ans: 232,
		},
		{
			N:   500,
			Ans: 124,
		},
	} {
		if ans := trailingZeroes(tc.N); ans != tc.Ans {
			t.Fatalf("n=%d, expect %d get %d", tc.N, tc.Ans, ans)
		}
	}
}
