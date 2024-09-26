package leetcode

import "testing"

func TestMaxDistToClosest(t *testing.T) {
	seats, exp := []int{1, 0, 0, 0, 1, 0, 1}, 2
	if r := maxDistToClosest(seats); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	seats, exp = []int{1, 0, 0, 0}, 3
	if r := maxDistToClosest(seats); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	seats, exp = []int{0, 1}, 1
	if r := maxDistToClosest(seats); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
