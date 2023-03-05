package leetcode

import (
	"reflect"
	"testing"
)

func TestSolove(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	exp := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	if !reflect.DeepEqual(exp, board) {
		t.Fatalf("expect %v get %v", exp, board)
	}
}
