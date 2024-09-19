package leetcode

import (
	"reflect"
	"testing"
)

func TestSmallestTrimmedNumbers(t *testing.T) {
	nums, queries, exp := []string{"102", "473", "251", "814"}, [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}, []int{2, 2, 1, 0}
	if r := smallestTrimmedNumbers(nums, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums, queries, exp = []string{"24", "37", "96", "04"}, [][]int{{2, 1}, {2, 2}}, []int{3, 0}
	if r := smallestTrimmedNumbers(nums, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
