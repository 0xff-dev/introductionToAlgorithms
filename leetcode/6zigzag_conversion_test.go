package leetcode

import "testing"

type ans struct {
	Input string
	Row   int
	Ans   string
}

func TestConvert(t *testing.T) {
	for idx, a := range []ans{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"", 1, ""},
		{"aaab", 2, "aaab"},
		{"abcdef", 1, "abcdef"},
	} {
		t.Logf("test count %d\n", idx)
		if r := convert(a.Input, a.Row); r != a.Ans {
			t.Fatalf("input[%s] expect to get[%s], result is[%s]", a.Input, a.Ans, r)
		}
		if r := convert2(a.Input, a.Row); r != a.Ans {
			t.Fatalf("input[%s] expect to get[%s], result is[%s]", a.Input, a.Ans, r)
		}
	}
}
