package leetcode

import "testing"

func TestCalPoints(t *testing.T) {
	ops := []string{"5", "2", "C", "D", "+"}
	if r := calPoints(ops); r != 30 {
		t.Fatalf("expect 30 get %d", r)
	}

	ops = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	if r := calPoints(ops); r != 27 {
		t.Fatalf("expect 27 get %d", r)
	}

	ops = []string{"1"}
	if r := calPoints(ops); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
