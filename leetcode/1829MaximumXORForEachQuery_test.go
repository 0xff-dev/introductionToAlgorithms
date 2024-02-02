package leetcode

import (
	"reflect"
	"testing"
)

func TestGetMaximumXor(t *testing.T) {
	nums := []int{0, 1, 1, 3}
	maxBit := 2
	exp := []int{0, 3, 2, 3}
	if r := getMaximumXor(nums, maxBit); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums = []int{2, 3, 4, 7}
	maxBit = 3
	exp = []int{5, 2, 6, 5}
	if r := getMaximumXor(nums, maxBit); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{0, 1, 2, 2, 5, 7}
	maxBit = 3
	exp = []int{4, 3, 6, 4, 6, 7}
	if r := getMaximumXor(nums, maxBit); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
