package leetcode

import "testing"

func TestFirstUniqChar(t *testing.T) {
	s := "leetcode"
	if r := firstUniqChar(s); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	s = "loveleetcode"
	if r := firstUniqChar(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "aabb"
	if r := firstUniqChar(s); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
