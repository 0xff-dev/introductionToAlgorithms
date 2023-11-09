package leetcode

import "testing"

func TestCountHomogenous(t *testing.T) {
	s := "abbcccaa"
	if r := countHomogenous(s); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}
	s = "xy"
	if r := countHomogenous(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	s = "zzzzz"
	if r := countHomogenous(s); r != 15 {
		t.Fatalf("expect 15 get %d", r)
	}
}
