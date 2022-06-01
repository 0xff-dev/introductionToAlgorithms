package leetcode

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := []int{1, 3, 6, 10}
	if r := runningSum(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	nums = []int{1, 1, 1, 1, 1}
	expect = []int{1, 2, 3, 4, 5}
	if r := runningSum(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}
