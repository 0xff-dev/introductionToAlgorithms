package leetcode

import "testing"

func TestMinCost1578(t *testing.T) {
	colors := "abaac"
	nt := []int{1, 2, 3, 4, 5}
	if r := minCost1578(colors, nt); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	colors = "abc"
	nt = []int{1, 2, 3}
	if r := minCost1578(colors, nt); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	colors = "aabaa"
	nt = []int{1, 2, 3, 4, 1}
	if r := minCost1578(colors, nt); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
