package leetcode

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	exp := []int{3, 3, 5, 5, 6, 7}
	if r := maxSlidingWindow(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
