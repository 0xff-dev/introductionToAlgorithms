package leetcode

import "testing"

func TestTotalNQueens(t *testing.T) {
	n := 4
	if r := totalNQueens(n); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
