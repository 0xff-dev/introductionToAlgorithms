package leetcode

import "testing"

func TestMaxCompatibilitySum(t *testing.T) {
	students, mentors, exp := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}, 8
	if r := maxCompatibilitySum(students, mentors); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	students, mentors, exp = [][]int{{0, 0}, {0, 0}, {0, 0}}, [][]int{{1, 1}, {1, 1}, {1, 1}}, 0
	if r := maxCompatibilitySum(students, mentors); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

}
