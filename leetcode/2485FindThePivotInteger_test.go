package leetcode

import "testing"

func TestPivotInteger(t *testing.T) {
	n := 8
	if r := pivotInteger(n); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	if r := pivotInteger(1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := pivotInteger(4); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
