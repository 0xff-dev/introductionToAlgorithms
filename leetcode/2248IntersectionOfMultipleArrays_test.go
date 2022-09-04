package leetcode

import (
	"reflect"
	"testing"
)

func TestIntersection2248(t *testing.T) {
	nums := [][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}
	exp := []int{3, 4}
	if r := intersection2248(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v ", exp, r)
	}

	nums = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	exp = []int{}
	if r := intersection2248(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
