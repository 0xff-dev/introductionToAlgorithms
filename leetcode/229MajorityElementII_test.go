package leetcode

import (
	"reflect"
	"testing"
)

func TestMajorityElement229(t *testing.T) {
	n := []int{3, 2, 3}
	exp := []int{3}
	if r := majorityElement229(n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = []int{1, 2}
	if r := majorityElement229(n); !reflect.DeepEqual(n, r) {
		t.Fatalf("expect %v get %v", n, r)
	}
}
