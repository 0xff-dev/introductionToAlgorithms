package leetcode

import "testing"

func TestContainsDupliate(t *testing.T) {
	n := []int{1, 2, 3, 1}
	if r := containsDuplicate(n); !r {
		t.Fatalf("expec true get false")
	}
	if r := containsDuplicate1(n); !r {
		t.Fatalf("expec true get false")
	}

	n = []int{1, 2, 3, 5}
	if r := containsDuplicate(n); r {
		t.Fatalf("expect false get true")
	}
	if r := containsDuplicate1(n); r {
		t.Fatalf("expec true get false")
	}

	n = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	if r := containsDuplicate(n); !r {
		t.Fatalf("expect true get false")
	}
	if r := containsDuplicate1(n); !r {
		t.Fatalf("expec true get false")
	}
}
