package leetcode

import "testing"

func TestChalkReplacer(t *testing.T) {
	chalk, k, exp := []int{5, 1, 5}, 22, 0
	if r := chalkReplacer(chalk, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	chalk, k, exp = []int{3, 4, 1, 2}, 25, 1
	if r := chalkReplacer(chalk, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
