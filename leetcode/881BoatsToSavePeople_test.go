package leetcode

import "testing"

func TestNumRescueBoats(t *testing.T) {
	people := []int{1, 2}
	limit := 3
	if r := numRescueBoats(people, limit); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	people = []int{3, 2, 2, 1}
	limit = 3
	if r := numRescueBoats(people, limit); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

