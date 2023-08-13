package leetcode

import "testing"

type input2369 struct {
	n      []int
	expect bool
}

func TestValidPartition(t *testing.T) {
	inputs := []input2369{
		{
			n:      []int{993335, 993336, 993337, 993338, 993339, 993340, 993341},
			expect: false,
		},
		{
			n:      []int{4, 4, 4, 5, 6},
			expect: true,
		},
		{
			n:      []int{1, 1, 1, 2},
			expect: false,
		},
	}
	for _, tc := range inputs {
		if r := validPartition(tc.n); r != tc.expect {
			t.Fatalf("with input %v, expect %v get %v", tc.n, tc.expect, r)
		}
	}
}
