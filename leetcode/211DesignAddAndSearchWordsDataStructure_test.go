package leetcode

import "testing"

func TestWordDictionary(t *testing.T) {
	c := Constructor211()
	c.AddWord("bad")
	c.AddWord("dad")
	c.AddWord("mad")

	if c.Search("pad") {
		t.Fatalf("search 'pad' expect false get true")
	}
	if !c.Search("bad") {
		t.Fatalf("search 'bad' expect true get false")
	}
	if !c.Search(".ad") {
		t.Fatalf("search '.ad' expect true get false")
	}
	if !c.Search("b..") {
		t.Fatalf("search 'b..' expect true get false")
	}
}
