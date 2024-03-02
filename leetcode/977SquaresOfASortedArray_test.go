package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	exp := []int{0, 1, 9, 16, 100}

	if r := sortedSquares(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{-7, -3, 2, 3, 11}
	exp = []int{4, 9, 9, 49, 121}
	if r := sortedSquares(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
