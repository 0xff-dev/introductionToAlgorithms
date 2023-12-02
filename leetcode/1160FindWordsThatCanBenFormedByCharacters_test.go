package leetcode

import "testing"

func TestCountCharacters(t *testing.T) {
	words := []string{"cat", "bt", "hat", "tree"}
	str := "atach"
	if r := countCharacters(words, str); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	words = []string{"hello", "world", "leetcode"}
	str = "welldonehoneyr"
	if r := countCharacters(words, str); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}
