package leetcode

import "testing"

func TestMctFromLeafValues(t *testing.T) {
	arr := []int{6, 2, 4}
	if r := mctFromLeafValues(arr); r != 32 {
		t.Fatalf("expect 32 get %d", r)
	}

	arr = []int{4, 11}
	if r := mctFromLeafValues(arr); r != 44 {
		t.Fatalf("expect 44 get %d", r)
	}
}
