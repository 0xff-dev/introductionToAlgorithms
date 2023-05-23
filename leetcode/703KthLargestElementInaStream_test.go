package leetcode

import "testing"

func TestContructor703(t *testing.T) {
	k := 3
	nums := []int{4, 5, 8, 2}
	c := Constructor703(k, nums)
	type testCase struct {
		input, got int
	}
	for _, tc := range []testCase{
		{3, 4}, {5, 5}, {10, 5}, {9, 8}, {4, 8},
	} {
		if r := c.Add(tc.input); r != tc.got {
			t.Fatalf("expect %d get %d", tc.got, r)
		}
	}
}
