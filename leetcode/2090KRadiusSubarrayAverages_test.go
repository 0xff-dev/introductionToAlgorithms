package leetcode

import (
	"reflect"
	"testing"
)

func TestGetAverage(t *testing.T) {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	exp := []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}
	if r := getAverages(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums = []int{1, 2, 3}
	k = 0
	if r := getAverages(nums, k); !reflect.DeepEqual(nums, r) {
		t.Fatalf("expect %v get %v", nums, r)
	}
	nums = []int{8}
	k = 10000
	exp = []int{-1}
	if r := getAverages(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{8}
	k = 1
	if r := getAverages(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
