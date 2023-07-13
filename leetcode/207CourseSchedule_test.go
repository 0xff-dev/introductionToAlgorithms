package leetcode

import "testing"

func TestCanFinish(t *testing.T) {
	n := 2
	p := [][]int{{1, 0}}
	if !canFinish(n, p) {
		t.Fatalf("expect true get false")
	}
	n = 2
	p = [][]int{{1, 0}, {0, 1}}
	if canFinish(n, p) {
		t.Fatalf("expect false get true")
	}
}
