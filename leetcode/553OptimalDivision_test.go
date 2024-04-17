package leetcode

import "testing"

func TestOptionalDivision(t *testing.T) {
	nums := []int{1000, 100, 10, 2}
	exp := "1000/(100/10/2)"
	if r := optimalDivision(nums); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	nums = []int{2, 3, 4}
	exp = "2/(3/4)"
	if r := optimalDivision(nums); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
