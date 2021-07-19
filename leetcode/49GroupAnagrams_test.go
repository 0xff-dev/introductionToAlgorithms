package leetcode

import "testing"

func TestGroupAnagrams(t *testing.T) {
	strs := []string{
		"eat", "tea", "tan",
		"ate", "nat", "bat",
	}
	t.Logf("%v", groupAnagrams(strs))

	strs = []string{""}
	t.Logf("%v", groupAnagrams(strs))

	strs = []string{"a"}
	t.Logf("%v", groupAnagrams(strs))
}
