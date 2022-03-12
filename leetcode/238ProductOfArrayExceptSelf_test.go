package leetcode

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := []int{24, 12, 8, 6}
	if r := productExceptSelf(nums); !reflect.DeepEqual(expect, r) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	nums = []int{-1, 1, 0, -3, 3}
	expect = []int{0, 0, 9, 0, 0}
	if r := productExceptSelf(nums); !reflect.DeepEqual(expect, r) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	nums = []int{2, 3, 5, 0}
	expect = []int{0, 0, 0, 30}
	if r := productExceptSelf(nums); !reflect.DeepEqual(expect, r) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}
