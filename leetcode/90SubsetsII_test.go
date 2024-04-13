package leetcode

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDuop(t *testing.T) {
	nums := []int{1, 2, 2}
	exp := [][]int{
		{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2},
	}
	if r := subsetsWithDup(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
