package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildArray1920(t *testing.T) {
	n := []int{0, 2, 1, 5, 3, 4}
	exp := []int{0, 1, 2, 4, 5, 3}
	if r := buildArray1920(n); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = []int{5, 0, 1, 2, 3, 4}
	exp = []int{4, 5, 0, 1, 2, 3}
	if r := buildArray1920(n); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
