package leetcode

import "testing"

type _testCase struct {
	N   int
	Ans int
}

func TestCountVowelStrings(t *testing.T) {
	for _, tstCast := range []_testCase{
		{1, 5},
		{2, 15},
		{33, 66045},
	} {
		if r := countVowelStrings(tstCast.N); r != tstCast.Ans {
			t.Fatalf("expect %d get %d", tstCast.Ans, r)
		}
	}
}
