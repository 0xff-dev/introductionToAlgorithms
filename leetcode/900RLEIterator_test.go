package leetcode

import "testing"

func TestConstructor900(t *testing.T) {
	encodings := []int{3, 8, 0, 9, 2, 5}
	input := []int{2, 1, 1, 2}
	exp := []int{8, 8, 5, -1}
	c := Constructor900(encodings)
	for i, n := range input {
		if r := c.Next(n); r != exp[i] {
			t.Fatalf("expect %d get %d", exp, r)
		}
	}
}
