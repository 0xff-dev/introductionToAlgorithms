package leetcode

import "testing"

func TestConstructor1032(t *testing.T) {
	words, q, exp := []string{"cd", "f", "kl"}, []byte("abcdefghijkl"), []bool{false, false, false, true, false, true, false, false, false, false, false, true}
	c := Constructor1032(words)
	for i := range q {
		if r := c.Query(q[i]); r != exp[i] {
			t.Fatalf("expect %v get %v", exp[i], r)
		}
	}
}
