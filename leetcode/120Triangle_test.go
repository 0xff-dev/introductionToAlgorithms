package leetcode

import "testing"

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	r := minimumTotal(triangle)
	if r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}
	if r := minimumTotal1(triangle); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}
	triangle = [][]int{
		{-10},
	}
	r = minimumTotal(triangle)
	if r != -10 {
		t.Fatalf("expect -10 get %d", r)
	}
	if r := minimumTotal1(triangle); r != -10 {
		t.Fatalf("expect -10 get %d", r)
	}
}
