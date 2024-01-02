package leetcode

import (
	"reflect"
	"testing"
)

func TestFindMatrix(t *testing.T) {
	nums := []int{1, 3, 4, 1, 2, 3, 1}
	exp := [][]int{
		{1, 2, 3, 4},
		{1, 3},
		{1},
	}
	if r := findMatrix(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{1, 2, 3, 4}
	exp = [][]int{
		{1, 2, 3, 4},
	}
	if r := findMatrix(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
