package leetcode

import "testing"

func TestNumFactoreBinaryTrees(t *testing.T) {
	arr := []int{2, 4}
	if r := numFactoredBinaryTrees(arr); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	arr = []int{2, 4, 5, 10}
	if r := numFactoredBinaryTrees(arr); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	arr = []int{2, 4, 5, 10, 100}
	if r := numFactoredBinaryTrees(arr); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}
}
