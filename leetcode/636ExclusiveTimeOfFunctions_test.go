package leetcode

import (
	"reflect"
	"testing"
)

func TestExclusiveTime(t *testing.T) {
	n := 2
	logs := []string{
		"0:start:0", "1:start:2", "1:end:5", "0:end:6",
	}
	exp := []int{3, 4}
	if r := exclusiveTime(n, logs); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
