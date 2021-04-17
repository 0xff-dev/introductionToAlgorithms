package leetcode

import "testing"

func TestIsHappy(t *testing.T) {
	if !isHappy(19) {
		t.Fatalf("19 is happy num")
	}

	if !isHappy(7) {
		t.Fatal("7 is happy num")
	}

	if isHappy(2) {
		t.Fatalf("2 isn't happy num")
	}
}
