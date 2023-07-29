package leetcode

import "testing"

func TestConstructor732(t *testing.T) {
	inputs := [][]int{
		{10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55},
	}
	exp := []int{1, 1, 2, 3, 3, 3}
	c := Constructor732()
	for idx, i := range inputs {
		if r := c.Book(i[0], i[1]); r != exp[idx] {
			t.Fatalf("expect %d get %d", exp[idx], r)
		}
	}
}
