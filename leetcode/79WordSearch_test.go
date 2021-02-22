package leetcode

import "testing"

func TestExist(t *testing.T) {
	input := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	if !exist(input, word) {
		t.Fatal("should return true")
	}

	input = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "SEE"
	if !exist(input, word) {
		t.Fatal("should return true")
	}

	input = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCB"
	if exist(input, word) {
		t.Fatal("should return false")
	}

	t.Log("test dbac")
	input = [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	word = "dbac"
	if !exist(input, word) {
		t.Fatal("should return ture")
	}
}
