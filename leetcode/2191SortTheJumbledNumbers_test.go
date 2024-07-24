package leetcode

import (
	"reflect"
	"testing"
)

func TestSortJumbled(t *testing.T) {
	mapping, nums, exp := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}, []int{338, 38, 991}
	if r := sortJumbled(mapping, nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	mapping, nums, exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}, []int{123, 456, 789}
	if r := sortJumbled(mapping, nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	mapping, nums, exp = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if r := sortJumbled(mapping, nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
