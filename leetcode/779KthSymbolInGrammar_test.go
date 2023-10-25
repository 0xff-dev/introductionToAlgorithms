package leetcode

import "testing"

func TestKthGrammar(t *testing.T) {
	n := 4
	exp := []int{0, 0, 1, 1, 0, 1, 0, 0, 1}
	for i := 1; i <= 8; i++ {
		if r := kthGrammar(n, i); r != exp[i] {
			t.Fatalf("expect %d get %d", exp[i], r)
		}
	}
}
