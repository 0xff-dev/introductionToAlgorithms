package leetcode

import (
	"reflect"
	"testing"
)

func TestFindArray(t *testing.T) {
	pref := []int{5, 2, 0, 3, 1}
	exp := []int{5, 7, 2, 3, 2}
	if r := findArray(pref); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
