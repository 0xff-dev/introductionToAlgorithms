package leetcode

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	nums := []int{1, 2, 1}
	expect := []int{1, 2, 1, 1, 2, 1}
	if r := getConcatenation(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	nums = []int{1, 3, 2, 1}
	expect = []int{1, 3, 2, 1, 1, 3, 2, 1}
	if r := getConcatenation(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}
