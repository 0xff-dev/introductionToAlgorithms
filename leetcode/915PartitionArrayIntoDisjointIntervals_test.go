package leetcode

import "testing"

func TestPartitionDisjoint(t *testing.T) {
	nums := []int{5, 0, 3, 8, 6}
	if r := partitionDisjoint(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{0, 3, 5, 8, 6}
	if r := partitionDisjoint(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{1, 2}
	if r := partitionDisjoint(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{1, 1, 1, 0, 6, 12}
	if r := partitionDisjoint(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
