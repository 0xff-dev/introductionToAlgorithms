package leetcode

import (
	"reflect"
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	a, b := []int{35, 17, 4, 24, 10}, []int{63, 21}
	exp := []int{24, 21}
	if r := fairCandySwap(a, b); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
