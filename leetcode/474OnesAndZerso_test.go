package leetcode

import "testing"

func TestFindMaxForm(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	if r := findMaxForm(strs, 5, 3); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	strs = []string{"10", "0", "1"}
	if r := findMaxForm(strs, 1, 1); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	strs = []string{"001", "110", "0000", "0000"}
	if r := findMaxForm(strs, 9, 2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	strs = []string{"10", "0001", "111001", "1", "0"}
	if r := findMaxForm(strs, 4, 3); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	strs = []string{"10", "1", "1", "0"}
	if r := findMaxForm(strs, 2, 2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	strs = []string{"00", "000"}
	if r := findMaxForm(strs, 1, 10); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	strs = []string{"11", "111"}
	if r := findMaxForm(strs, 10, 1); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	strs = []string{"0", "0", "1", "1"}
	if r := findMaxForm(strs, 2, 2); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	strs = []string{"00101011"}
	if r := findMaxForm(strs, 36, 39); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
