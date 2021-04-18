package leetcode

import "testing"

func TestArrayNesting(t *testing.T) {
	input := []int{5,4,0,3,1,6,2}
	if r := arrayNesting(input); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
