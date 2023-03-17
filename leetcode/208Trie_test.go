package leetcode

import "testing"

func TestTrie(t *testing.T) {
	c := Constructor208()
	c.Insert("apple")
	if !c.Search("apple") {
		t.Fatalf("search 'apple' expect true get false")
	}

	if c.Search("app") {
		t.Fatalf("search 'app' expect false get true")
	}

	if !c.StartsWith("app") {
		t.Fatalf("startWith 'app' expect true get false")
	}

	c.Insert("app")
	if !c.Search("app") {
		t.Fatalf("search 'app' expect true get false")
	}
}
