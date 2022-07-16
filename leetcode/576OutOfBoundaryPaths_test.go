package leetcode

import "testing"

func TestFindPaths(t *testing.T) {
	if r := findPaths(2, 2, 2, 0, 0); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	if r := findPaths(1, 3, 3, 0, 1); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	if r := findPaths(8, 7, 16, 1, 5); r != 102984580 {
		t.Fatalf("expect 102984580 get %d", r)
	}
}
