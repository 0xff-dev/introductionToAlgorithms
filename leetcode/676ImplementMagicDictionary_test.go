package leetcode

import "testing"

func TestConstructor676(t *testing.T) {
	c := Constructor676()
	c.BuildDict([]string{"hello", "leetcode"})
	type input struct {
		searchWord string
		exp        bool
	}
	for _, tc := range []input{
		{"hello", false}, {"hhllo", true}, {"hell", false}, {"leetcoded", false},
	} {
		if r := c.Search(tc.searchWord); r != tc.exp {
			t.Fatalf("with input %s expect %v get %v", tc.searchWord, tc.exp, r)
		}
	}
}
