package leetcode

import "testing"

type numberOfStepsInput struct {
	input int
	ans   int
}

func TestNumberOfSteps(t *testing.T) {
	for _, tc := range []numberOfStepsInput{
		{
			5, 4,
		},
		{
			14, 6,
		},
		{
			8, 4,
		},
		{
			123, 12,
		},
	} {
		if r := numberOfSteps(tc.input); r != tc.ans {
			t.Fatalf("expect %d get %d", tc.ans, r)
		}
	}
}
