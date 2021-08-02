package leetcode

import "testing"

func TestRemoveDuplicatesII(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	if r := removeDuplicatesII(nums); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	t.Logf("%v", nums)

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	if r := removeDuplicatesII(nums); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	t.Logf("%v", nums)

	nums = []int{1}
	if r := removeDuplicatesII(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	t.Logf("%v", nums)

	nums = []int{1, 1, 1, 1, 1}
	if r := removeDuplicatesII(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	t.Logf("%v", nums)

	nums = []int{1, 1, 1, 1, 2, 2, 2, 3, 4, 4}
	if r := removeDuplicatesII(nums); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	t.Logf("%v", nums)

	nums = []int{1, 1, 2}
	if r := removeDuplicatesII(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	t.Logf("%v", nums)
}
