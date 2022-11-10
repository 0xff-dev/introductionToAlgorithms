package leetcode

import "testing"

func TestRemoveDuplicates1047(t *testing.T) {
	s := "abbaca"
	if r := removeDuplicates1047(s); r != "ca" {
		t.Fatalf("expect 'ca' get %s", r)
	}

	s = "azxxzy"
	if r := removeDuplicates1047(s); r != "ay" {
		t.Fatalf("expect 'ay' get %s", r)
	}
}
