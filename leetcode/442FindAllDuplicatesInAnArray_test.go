package leetcode

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	exp := []int{2, 3}
	if r := findDuplicates(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{1, 1, 2}
	exp = []int{1}
	if r := findDuplicates(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{1}
	exp = []int{}
	if r := findDuplicates(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
