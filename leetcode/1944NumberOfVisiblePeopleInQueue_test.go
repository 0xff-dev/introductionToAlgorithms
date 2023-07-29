package leetcode

import (
	"reflect"
	"testing"
)

func TestCanSeePersonsCount(t *testing.T) {
	h := []int{10, 6, 8, 5, 11, 9}
	exp := []int{3, 1, 2, 1, 1, 0}
	if r := canSeePersonsCount(h); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	h = []int{5, 1, 2, 3, 10}
	exp = []int{4, 1, 1, 1, 0}
	if r := canSeePersonsCount(h); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
