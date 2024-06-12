package leetcode

import "testing"

func TestNumTeams(t *testing.T) {
	rating := []int{2, 5, 3, 4, 1}
	exp := 3
	if r := numTeams(rating); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	rating = []int{2, 1, 3}
	exp = 0
	if r := numTeams(rating); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	rating = []int{1, 2, 3, 4}
	exp = 4
	if r := numTeams(rating); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
