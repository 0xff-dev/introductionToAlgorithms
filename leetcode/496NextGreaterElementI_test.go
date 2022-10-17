package leetcode

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	n1, n2 := []int{4, 1, 2}, []int{1, 3, 4, 2}
	exp := []int{-1, 3, -1}
	if r := nextGreaterElement(n1, n2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n1, n2 = []int{2, 4}, []int{1, 2, 3, 4}
	exp = []int{3, -1}
	if r := nextGreaterElement(n1, n2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

}
