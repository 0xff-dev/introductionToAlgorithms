package leetcode

import "testing"

func TestContructor731(t *testing.T) {
	inputs := [][]int{
		{10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55},
	}
	exp := []bool{true, true, true, false, true, true}
	c := Constructor731()
	for idx, in := range inputs {
		if r := c.Book(in[0], in[1]); r != exp[idx] {
			t.Fatalf("with input %v expect %v get %v", in, exp[idx], r)
		}
	}
}
