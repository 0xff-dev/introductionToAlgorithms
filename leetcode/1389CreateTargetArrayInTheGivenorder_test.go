package leetcode

import (
	"reflect"
	"testing"
)

func TestCreateTargetArray(t *testing.T) {
	n := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	exp := []int{0, 4, 1, 3, 2}
	if r := createTargetArray(n, index); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n = []int{1, 2, 3, 4, 0}
	index = []int{0, 1, 2, 3, 0}
	exp = []int{0, 1, 2, 3, 4}
	if r := createTargetArray(n, index); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
