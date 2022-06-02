package leetcode

import (
	"reflect"
	"testing"
)

type input119 struct {
	input  int
	expect []int
}

func TestGetRow(t *testing.T) {
	for _, tc := range []input119{
		{
			input:  3,
			expect: []int{1, 3, 3, 1},
		}, {
			input:  0,
			expect: []int{1},
		}, {
			input:  1,
			expect: []int{1, 1},
		}, {
			input:  2,
			expect: []int{1, 2, 1},
		},
	} {
		if r := getRow(tc.input); !reflect.DeepEqual(r, tc.expect) {
			t.Fatalf("expect %v get %v", tc.expect, r)
		}
	}
}
