package leetcode

import (
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	n1, n2 := []int{1, 2, 3}, []int{2, 4, 6}
	exp := [][]int{
		{3, 1}, {6, 4},
	}
	if r := findDifference(n1, n2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
