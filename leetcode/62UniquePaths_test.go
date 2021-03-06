package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	m, n := 3, 7
	if r := uniquePaths(m, n); r != 28 {
		t.Fatalf("expect 28 get %d", r)
	}

	m, n = 3, 2
	if r := uniquePaths(m, n); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	m, n = 7, 3
	if r := uniquePaths(m, n); r != 28 {
		t.Fatalf("expect 28 get %d", r)
	}

	m, n = 3, 3
	if r := uniquePaths(m, n); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
