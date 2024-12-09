package leetcode

import (
	"reflect"
	"testing"
)

func TestIsArraySpecial(t *testing.T) {
	nums, queries, exp := []int{3, 4, 1, 2, 6}, [][]int{{0, 4}}, []bool{false}
	if r := isArraySpecial(nums, queries); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums, queries, exp = []int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}, []bool{false, true}
	if r := isArraySpecial(nums, queries); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
