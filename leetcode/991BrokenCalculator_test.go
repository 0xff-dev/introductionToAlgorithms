package leetcode

import "testing"

func TestBrokenCalc(t *testing.T) {
	start, target := 2, 3
	if r := brokenCalc(start, target); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	start, target = 5, 8
	if r := brokenCalc(start, target); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	start, target = 3, 10
	if r := brokenCalc(start, target); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
