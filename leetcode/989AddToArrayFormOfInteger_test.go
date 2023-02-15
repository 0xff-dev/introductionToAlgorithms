package leetcode

import (
	"reflect"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	num, k := []int{1, 2, 0, 0}, 34
	exp := []int{1, 2, 3, 4}
	if r := addToArrayForm(num, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	num, k = []int{2, 7, 4}, 181
	exp = []int{4, 5, 5}
	if r := addToArrayForm(num, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	num, k = []int{2, 1, 5}, 806
	exp = []int{1, 0, 2, 1}
	if r := addToArrayForm(num, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
