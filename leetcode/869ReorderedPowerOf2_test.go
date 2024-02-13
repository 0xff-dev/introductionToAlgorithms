package leetcode

import "testing"

func TestReorderedPowerOf2(t *testing.T) {
	n := 46
	exp := true
	if r := reorderedPowerOf2(n); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 125
	exp = true
	if r := reorderedPowerOf2(n); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 1
	exp = true
	if r := reorderedPowerOf2(n); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 10
	exp = false
	if r := reorderedPowerOf2(n); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
