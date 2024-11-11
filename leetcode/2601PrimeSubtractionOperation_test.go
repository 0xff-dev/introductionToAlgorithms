package leetcode

import "testing"

func TestPrimeSubOperation(t *testing.T) {
	nums, exp := []int{4, 9, 6, 10}, true
	if r := primeSubOperation(nums); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
