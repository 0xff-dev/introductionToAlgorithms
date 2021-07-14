package leetcode

import "testing"

func TestIsIsomorphic(t *testing.T) {
	s, t1 := "egg", "add"
	if r := isIsomorphic(s, t1); !r {
		t.Fatalf("expect true get false")
	}

	s, t1 = "foo", "bar"
	if r := isIsomorphic(s, t1); r {
		t.Fatalf("expect false get true")
	}

	s, t1 = "bar", "foo"
	if r := isIsomorphic(s, t1); r {
		t.Fatalf("expect false get true")
	}

	s, t1 = "paper", "title"
	if r := isIsomorphic(s, t1); !r {
		t.Fatalf("expect true get false")
	}

	s, t1 = "abcd", "efgh"
	if r := isIsomorphic(s, t1); !r {
		t.Fatalf("expect true get false")
	}
}
