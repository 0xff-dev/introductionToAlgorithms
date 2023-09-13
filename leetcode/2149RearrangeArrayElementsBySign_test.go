package leetcode

import (
	"reflect"
	"testing"
)

func TestRearrangeArray(t *testing.T) {
	in := []int{3, 1, -2, -5, 2, -4}
	exp := []int{3, -2, 1, -5, 2, -4}
	if r := rearrangeArray(in); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
