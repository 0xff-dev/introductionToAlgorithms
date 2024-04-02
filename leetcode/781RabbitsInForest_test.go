package leetcode

import "testing"

func TestNumrabbits(t *testing.T) {
	answers := []int{1, 1, 2}
	exp := 5
	if r := numRabbits(answers); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	answers = []int{10, 10, 10}
	exp = 11
	if r := numRabbits(answers); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
