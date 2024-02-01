package leetcode

import (
	"reflect"
	"testing"
)

func TestDivideArray2966(t *testing.T) {
	nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
	k := 2
	exp := [][]int{
		{1, 1, 3}, {3, 4, 5}, {7, 8, 9},
	}
	if r := divideArray2966(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{1, 3, 3, 2, 7, 3}
	k = 3
	if r := divideArray2966(nums, k); r != nil {
		t.Fatalf("expect %v get %v", exp, r)
	}

}
