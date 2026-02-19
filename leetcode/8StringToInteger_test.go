package leetcode

import "testing"

type input8 struct {
	s   string
	exp int
}

func TestNyAtoi(t *testing.T) {
	inputs := []input8{
		{
			"42", 42,
		},
		{
			"-042", -42,
		},
		{
			"1337c0d3", 1337,
		},
		{
			"0-1", 0,
		},
		{
			"words and 987", 0,
		},
		{
			"-2147483647", -2147483647,
		},
		{
			"1095502006p8", 1095502006,
		},
	}

	for i := range inputs {
		got := myAtoi(inputs[i].s)
		if got != inputs[i].exp {
			t.Fatalf("except %d got %d", inputs[i].exp, got)
		}
	}
}
