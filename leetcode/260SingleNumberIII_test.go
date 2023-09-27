package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestSingleNumber260(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2, 5}
	exp := []int{3, 5}
	r := singleNumber260(nums)
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{0, 1}
	r = singleNumber260(nums)
	exp = []int{0, 1}
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
