package leetcode

import (
	"reflect"
	"testing"
)

func TestUniquePattern(t *testing.T) {
	str := "abb"
	exp := []int{1, 2, 2}
	if r := uniquePattern(str); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	str = "a"
	exp = []int{1}
	if r := uniquePattern(str); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

func TestFindAndReplacePattern(t *testing.T) {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	exp := []string{"mee", "aqq"}
	if r := findAndReplacePattern(words, pattern); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
