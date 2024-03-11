package leetcode

import "testing"

func TestLeastInterval(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	exp := 8
	if r := leastInterval(tasks, n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	tasks = []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	n = 1
	exp = 6
	if r := leastInterval(tasks, n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	tasks = []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n = 2
	exp = 16
	if r := leastInterval(tasks, n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
