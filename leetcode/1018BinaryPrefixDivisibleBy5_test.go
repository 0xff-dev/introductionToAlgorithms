package leetcode

import (
	"reflect"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	nums := []int{0, 1, 1}
	expect := []bool{true, false, false}
	if r := prefixesDivBy5(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	nums = []int{1, 1, 1}
	expect = []bool{false, false, false}
	if r := prefixesDivBy5(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
	nums = []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	expect = []bool{false, false, false, false, false, true, false, false, false, true, false, false, true, false, false, false, false, true, true}
	if r := prefixesDivBy5(nums); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}
