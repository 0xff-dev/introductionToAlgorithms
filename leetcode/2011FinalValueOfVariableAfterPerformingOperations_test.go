package leetcode

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	operations := []string{"--X", "X++", "X++"}
	if r := finalValueAfterOperations(operations); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	operations = []string{"++X", "++X", "X++"}
	if r := finalValueAfterOperations(operations); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

}
