package leetcode

import "testing"

func TestFindDifferentBinaryString(t *testing.T) {
	nums := []string{"01", "10"}
	if r := findDifferentBinaryString(nums); r != "00" {
		t.Fatalf("expect %s get %s", "00", r)
	}
}
