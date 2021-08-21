package leetcode

import (
	"reflect"
	"testing"
)

func TestRoate(t *testing.T) {
	nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3
	expect := []int{5, 6, 7, 1, 2, 3, 4}
	if rotate1(nums, k); !reflect.DeepEqual(nums, expect) {
		t.Fatalf("expect %v get %v", expect, nums)
	}
}
