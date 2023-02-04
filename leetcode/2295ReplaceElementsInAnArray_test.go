package leetcode

import (
	"reflect"
	"testing"
)

func TestArrayChange(t *testing.T) {
	nums, ops := []int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}
	exp := []int{3, 2, 7, 1}
	if r := arrayChange(nums, ops); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums, ops = []int{1, 2}, [][]int{{1, 3}, {2, 1}, {3, 2}}
	exp = []int{2, 1}
	if r := arrayChange(nums, ops); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
