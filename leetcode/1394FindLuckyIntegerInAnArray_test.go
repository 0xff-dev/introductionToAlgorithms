package leetcode

import "testing"

func TestFindLucky(t *testing.T) {
	arr := []int{2, 2, 3, 4}
	if r := findLucky(arr); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	arr = []int{1, 2, 2, 3, 3, 3}
	if r := findLucky(arr); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	arr = []int{2, 2, 2, 3, 3}
	if r := findLucky(arr); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
