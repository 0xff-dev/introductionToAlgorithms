package leetcode

import (
	"reflect"
	"testing"
)

func TestMaxSubsequence(t *testing.T) {
	nums, k, exp := []int{2, 1, 3, 3}, 2, []int{3, 3}
	if r := maxSubsequence(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums, k, exp = []int{-1, -2, 3, 4}, 3, []int{-1, 3, 4}
	if r := maxSubsequence(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums, k, exp = []int{3, 4, 3, 3}, 2, []int{3, 4}
	if r := maxSubsequence(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
