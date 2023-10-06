package leetcode

import (
	"reflect"
	"testing"
)

func TestExecuteInstructions(t *testing.T) {
	n := 3
	pos := []int{0, 1}
	s := "RRDDLU"
	exp := []int{1, 5, 4, 3, 1, 0}
	if r := executeInstructions(n, pos, s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 2
	pos = []int{1, 1}
	s = "LURD"
	exp = []int{4, 1, 0, 0}
	if r := executeInstructions(n, pos, s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 1
	pos = []int{0, 0}
	s = "LRUD"
	exp = []int{0, 0, 0, 0}
	if r := executeInstructions(n, pos, s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
