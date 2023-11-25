package leetcode

import (
	"reflect"
	"testing"
)

func TestGetSumAbsoluteDifference(t *testing.T) {
	nums := []int{2, 3, 5}
	exp := []int{4, 3, 5}
	if r := getSumAbsoluteDifferences(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{1, 4, 6, 8, 10}
	exp = []int{24, 15, 13, 15, 21}
	if r := getSumAbsoluteDifferences(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
