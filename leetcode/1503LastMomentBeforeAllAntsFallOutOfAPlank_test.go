package leetcode

import "testing"

func TestGetLastMomnet(t *testing.T) {
	n := 4
	left := []int{4, 3}
	right := []int{0, 1}
	if r := getLastMoment(n, left, right); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	n = 7
	left = []int{}
	right = []int{0, 1, 2, 3, 4, 5, 6, 7}
	if r := getLastMoment(n, left, right); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
