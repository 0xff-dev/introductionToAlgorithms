package leetcode

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3}
	exp := []int{2, 3, 4, -1, 4}
	if r := nextGreaterElements(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{5, 4, 3, 2, 1}
	exp = []int{-1, 5, 5, 5, 5}
	if r := nextGreaterElements(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{1, 2, 1}
	exp = []int{2, -1, 2}
	if r := nextGreaterElements(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{1, 2}
	exp = []int{2, -1}
	if r := nextGreaterElements(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums = []int{2, 1}
	exp = []int{-1, 2}
	if r := nextGreaterElements(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
