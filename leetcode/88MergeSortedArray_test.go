package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge88(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge88(nums1, 3, nums2, 3)
	expect := []int{1, 2, 2, 3, 5, 6}
	if !reflect.DeepEqual(nums1, expect) {
		t.Fatalf("expect %v get %v", expect, nums1)
	}

	nums1, nums2 = []int{1}, []int{}
	merge88(nums1, 1, nums2, 0)
	expect = []int{1}
	if !reflect.DeepEqual(nums1, expect) {
		t.Fatalf("expect %v get %v", expect, nums1)
	}

	nums1, nums2 = []int{0}, []int{1}
	merge88(nums1, 0, nums2, 1)
	expect = []int{1}
	if !reflect.DeepEqual(nums1, expect) {
		t.Fatalf("expect %v get %v", expect, nums1)
	}
}
