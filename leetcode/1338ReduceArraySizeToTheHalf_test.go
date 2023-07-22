package leetcode

import "testing"

func TestMinSetSize(t *testing.T) {
	arr := []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
	if r := minSetSize(arr); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	arr = []int{7, 7, 7, 7, 7, 7}
	if r := minSetSize(arr); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
