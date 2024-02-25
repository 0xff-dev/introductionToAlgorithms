package leetcode

import "testing"

func TestCanTraverseAllPairs(t *testing.T) {
	nums := []int{2, 3, 6}
	if r := canTraverseAllPairs(nums); !r {
		t.Fatalf("expect true get false")
	}

	nums = []int{3, 9, 5}
	if r := canTraverseAllPairs(nums); r {
		t.Fatalf("expect false get true")
	}
	nums = []int{4, 3, 12, 8}
	if r := canTraverseAllPairs(nums); !r {
		t.Fatalf("expect true get false")
	}

}
