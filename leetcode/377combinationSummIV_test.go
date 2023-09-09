package leetcode

import "testing"

func TestCombinationSum4(t *testing.T) {
	n := []int{1, 2, 3}
	target := 4
	if r := combinationSum4(n, target); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	n = []int{9}
	target = 3
	if r := combinationSum4(n, target); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	n = []int{2, 1, 3}
	target = 35
	if r := combinationSum4(n, target); r != 1132436852 {
		t.Fatalf("expect 1132436852 get %d", r)
	}
}
