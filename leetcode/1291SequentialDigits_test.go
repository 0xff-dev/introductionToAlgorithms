package leetcode

import (
	"reflect"
	"testing"
)

func TestSequentialDigits(t *testing.T) {
	l, h := 100, 300
	exp := []int{123, 234}
	if r := sequentialDigits(l, h); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	l, h = 1000, 13000
	exp = []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}
	if r := sequentialDigits(l, h); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
