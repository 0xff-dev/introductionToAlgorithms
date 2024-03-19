package leetcode

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	bits := []int{1, 0, 0}
	exp := true
	if r := isOneBitCharacter(bits); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	bits = []int{1, 1, 1, 0}
	exp = false
	if r := isOneBitCharacter(bits); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
