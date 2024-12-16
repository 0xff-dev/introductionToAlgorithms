package leetcode

import (
	"reflect"
	"testing"
)

func TestGetFinalState(t *testing.T) {
	nums, k, multiplier, exp := []int{2, 1, 3, 5, 6}, 5, 2, []int{8, 4, 6, 5, 6}
	if r := getFinalState(nums, k, multiplier); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums, k, multiplier, exp = []int{1, 2}, 3, 4, []int{16, 8}
	if r := getFinalState(nums, k, multiplier); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
