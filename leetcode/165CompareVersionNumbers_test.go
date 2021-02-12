package leetcode

import "testing"

type compareType struct {
	V1, V2 string
	Ans    int
}

func TestCompareVersion(t *testing.T) {
	for _, testCase := range []compareType{
		{
			"1.01", "1.001", 0,
		},
		{
			"1.0", "1.00", 0,
		},
		{
			"0.1", "1.1", -1,
		},
		{
			"1.0.1", "1", 1,
		},
		{
			"7.5.2.4", "7.5.3", -1,
		},
		{
			"1", "1.1", -1,
		},
	} {
		if r := compareVersion(testCase.V1, testCase.V2); r != testCase.Ans {
			t.Fatalf("v1[%s] v2[%s] ans[%d], func result[%d]", testCase.V1, testCase.V2, testCase.Ans, r)
		}
	}
}
