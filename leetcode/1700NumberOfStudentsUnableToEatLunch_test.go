package leetcode

import "testing"

func TestCountStudents(t *testing.T) {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	exp := 0
	if r := countStudents(students, sandwiches); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	students = []int{1, 1, 1, 0, 0, 1}
	sandwiches = []int{1, 0, 0, 0, 1, 1}
	exp = 3
	if r := countStudents(students, sandwiches); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
