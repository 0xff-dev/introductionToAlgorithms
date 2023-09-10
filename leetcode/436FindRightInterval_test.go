package leetcode

import (
	"reflect"
	"testing"
)

func TestFindRightInterval(t *testing.T) {
	i := [][]int{{1, 2}}
	exp := []int{-1}
	if r := findRightInterval(i); !reflect.DeepEqual(exp, r) {
		t.Fatalf("with %v expect %v get %v", i, exp, r)
	}
	i = [][]int{{3, 4}, {2, 3}, {1, 2}}
	exp = []int{-1, 0, 1}
	if r := findRightInterval(i); !reflect.DeepEqual(exp, r) {
		t.Fatalf("with %v expect %v get %v", i, exp, r)
	}
	i = [][]int{{1, 4}, {2, 3}, {3, 4}}
	exp = []int{-1, 2, -1}
	if r := findRightInterval(i); !reflect.DeepEqual(exp, r) {
		t.Fatalf("with %v expect %v get %v", i, exp, r)
	}
	i = [][]int{{1, 1}, {3, 4}}
	exp = []int{0, -1}
	if r := findRightInterval(i); !reflect.DeepEqual(exp, r) {
		t.Fatalf("with %v expect %v get %v", i, exp, r)
	}
}
