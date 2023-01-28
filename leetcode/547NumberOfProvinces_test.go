package leetcode

import "testing"

func TestFindCircleNum(t *testing.T) {
	isConnected := [][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	}
	if r := findCircleNum(isConnected); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
