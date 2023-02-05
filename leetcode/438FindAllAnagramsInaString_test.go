package leetcode

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s, p := "cbaebabacd", "abc"
	exp := []int{0, 6}
	if r := findAnagrams(s, p); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	s, p = "abab", "ab"
	exp = []int{0, 1, 2}
	if r := findAnagrams(s, p); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
