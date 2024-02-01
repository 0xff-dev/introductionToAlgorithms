package leetcode

import "testing"

func TestNumOfSubarrays1524(t *testing.T) {
	arr := []int{1, 3, 5}
	exp := 4
	if r := numOfSubarrays1524(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	arr = []int{2, 4, 6}
	exp = 0
	if r := numOfSubarrays1524(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	arr = []int{1, 2, 3, 4, 5, 6, 7}
	exp = 16
	if r := numOfSubarrays1524(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
