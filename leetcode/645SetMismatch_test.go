package leetcode

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	nums := []int{1, 2, 2, 4}
	exp := []int{2, 3}
	if r := findErrorNums(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums = []int{1, 1}
	exp = []int{1, 2}
	if r := findErrorNums(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
