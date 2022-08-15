package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoOutOfThree(t *testing.T) {
	n1, n2, n3 := []int{1, 1, 3, 2}, []int{2, 3}, []int{3}
	exp := []int{2, 3}
	if r := twoOutOfThree(n1, n2, n3); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n1, n2, n3 = []int{3, 1}, []int{2, 3}, []int{1, 2}
	exp = []int{1, 2, 3}
	if r := twoOutOfThree(n1, n2, n3); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n1, n2, n3 = []int{1, 2, 2}, []int{4, 3, 3}, []int{5}
	exp = []int{}
	if r := twoOutOfThree(n1, n2, n3); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
