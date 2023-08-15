package leetcode

import (
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	nc := 2
	p := [][]int{{1, 0}}
	exp := []int{0, 1}
	if r := findOrder(nc, p); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nc = 1
	exp = []int{0}
	if r := findOrder(nc, nil); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
