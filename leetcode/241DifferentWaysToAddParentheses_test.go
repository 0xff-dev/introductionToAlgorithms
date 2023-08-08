package leetcode

import (
	"reflect"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	e := "2-1-1"
	exp := []int{0, 2}
	if r := diffWaysToCompute241(e); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	e = "2*3-4*5"
	exp = []int{-34, -14, -10, -10, 10}
	if r := diffWaysToCompute241(e); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
