package leetcode

import "testing"

func TestFindRotateSteps(t *testing.T) {
	ring := "godding"
	key := "gd"
	exp := 4
	if r := findRotateSteps(ring, key); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	key = "godding"
	exp = 13
	if r := findRotateSteps(ring, key); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ring = "abcde"
	key = "ade"
	exp = 6
	if r := findRotateSteps(ring, key); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ring = "pqwcx"
	key = "cpqwx"
	exp = 13
	if r := findRotateSteps(ring, key); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
