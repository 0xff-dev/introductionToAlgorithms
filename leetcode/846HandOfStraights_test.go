package leetcode

import "testing"

func TestIsNStraightHand(t *testing.T) {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	g := 3
	exp := true
	if r := isNStraightHand(hand, g); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	hand = []int{1, 2, 3, 4, 5}
	g = 4
	exp = false
	if r := isNStraightHand(hand, g); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}

	hand = []int{1, 1, 2, 2, 3, 3}
	g = 2
	exp = false
	if r := isNStraightHand(hand, g); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
