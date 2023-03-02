package leetcode

import "testing"

type compressType struct {
	Input []byte
	Ans   int
}

func TestCompress(t *testing.T) {
	for _, testCase := range []compressType{
		{
			[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			6,
		},
		{
			[]byte{'a'},
			1,
		},
		{
			[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			4,
		},
		{
			[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			6,
		},
		{
			[]byte{'a'},
			1,
		},
		{
			[]byte{'a', 'a'},
			2,
		},
	} {
		// if r := compress(testCase.Input); r != testCase.Ans {
		// 	t.Fatalf("input[%s] ans[%d] result[%d]", testCase.Input, testCase.Ans, r)
		// }
		if r := myCompress(testCase.Input); r != testCase.Ans {
			t.Fatalf("input[%s] ans[%d] result[%d]", testCase.Input, testCase.Ans, r)
		}
	}
}
