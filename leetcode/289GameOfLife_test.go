package leetcode

import (
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	board := [][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	}
	exp := [][]int{
		{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0},
	}
	gameOfLife(board)
	if !reflect.DeepEqual(exp, board) {
		t.Fatalf("expect %v get %v", exp, board)
	}
}
