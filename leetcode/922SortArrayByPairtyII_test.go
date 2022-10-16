package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArayByPairtyII(t *testing.T) {
	nums := []int{4, 2, 5, 7}
	exp := []int{4, 5, 2, 7}
	if r := sortArrayByParityII(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums = []int{2, 3}
	exp = []int{2, 3}
	if r := sortArrayByParityII(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
