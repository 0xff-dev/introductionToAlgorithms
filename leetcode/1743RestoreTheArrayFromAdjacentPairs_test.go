package leetcode

import (
	"reflect"
	"testing"
)

func TestRestoreArray(t *testing.T) {
	nums := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	exp := []int{1, 2, 3, 4, 5}
	exp1 := []int{5, 4, 3, 2, 1}
	if r := restoreArray(nums); !reflect.DeepEqual(exp, r) && !reflect.DeepEqual(exp1, r) {
		t.Fatalf("expect %v or %v get %v", exp, exp1, r)
	}
	nums = [][]int{{-2, 4}, {1, 4}, {-3, 1}}
	exp = []int{-2, 4, 1, -3}
	exp1 = []int{-3, 1, 4, -2}
	if r := restoreArray(nums); !reflect.DeepEqual(exp, r) && !reflect.DeepEqual(exp1, r) {
		t.Fatalf("expect %v or %v get %v", exp, exp1, r)
	}
}
