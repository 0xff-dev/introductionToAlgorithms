package leetcode

import "testing"

func TestPossibleBipartition(t *testing.T) {
	n, dislikes := 4, [][]int{{1, 2}, {1, 3}, {2, 4}}
	if !possibleBipartition(n, dislikes) {
		t.Fatalf("expect true get false")
	}

	n, dislikes = 3, [][]int{{1, 2}, {1, 3}, {2, 3}}
	if possibleBipartition(n, dislikes) {
		t.Fatalf("expect false get true")
	}

	n, dislikes = 5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	if possibleBipartition(n, dislikes) {
		t.Fatalf("expect false get true")
	}
}
