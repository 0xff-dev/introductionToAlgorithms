package leetcode

import (
	"testing"
)

func TestConstrucotr895(t *testing.T) {
	v := []int{5, 7, 5, 7, 4, 5}
	c := Constructor895()
	for _, item := range v {
		c.Push(item)
	}
	exp := []int{5, 7, 5, 4}
	for i := 0; i < 4; i++ {
		if r := c.Pop(); r != exp[i] {
			t.Fatalf("%d expect %d get %d", i, exp[i], r)
		}
	}
}
