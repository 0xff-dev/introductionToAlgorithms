package leetcode

import (
	"reflect"
	"testing"
)

func TestSub522(t *testing.T) {
	s := "abc"
	r := make([]string, 0)
	e := make(map[string]struct{})
	for l := len(s); l > 0; l-- {
		sub522(s, l, 0, e, []byte{}, &r)
	}
	exp := []string{"abc", "ab", "ac", "bc", "a", "b", "c"}
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

func TestFindLUSLength522(t *testing.T) {
	strs := []string{"abc", "cdc", "eae"}
	if r := findLUSlength522(strs); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	strs = []string{"aaa", "aaa", "aa"}
	if r := findLUSlength522(strs); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	strs = []string{"aabbcc", "aabbcc", "cb"}
	if r := findLUSlength522(strs); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
