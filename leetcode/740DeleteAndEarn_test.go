package leetcode

import "testing"

func TestDeleteAndEarn(t *testing.T) {
	input := []int{3, 2, 4}
	if r := deleteAndEarn(input); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	input = []int{2, 2, 3, 3, 3, 4}
	if r := deleteAndEarn(input); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
}
