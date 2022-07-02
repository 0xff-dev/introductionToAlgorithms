package leetcode

import (
	"reflect"
	"sort"
	"testing"
)


func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := [][]int{
		{},
		{1},
		{2},
		{3},
		{1, 2},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}

	r := subsets(nums)
	sort.Slice(r, func(i, j int)bool {
		return len(r[i]) < len(r[j])
	})
	for idx := 0; idx < len(r); idx++ {
		sort.Ints(r[idx])
	}
	if r := subsets(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}
