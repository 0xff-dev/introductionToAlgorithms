package leetcode

import (
	"reflect"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {
	s, p := []int{5, 1, 3}, []int{1, 2, 3, 4, 5}
	exp := []int{4, 0, 3}
	if r := successfulPairs(s, p, 7); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	s, p = []int{3, 1, 2}, []int{8, 5, 8}
	exp = []int{2, 0, 2}
	if r := successfulPairs(s, p, 16); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
