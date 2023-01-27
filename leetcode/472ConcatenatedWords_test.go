package leetcode

import (
	"reflect"
	"testing"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	words := []string{
		"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat",
	}
	exp := []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}
	if r := findAllConcatenatedWordsInADict(words); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	words = []string{"g", "e", "ge"}
	exp = []string{"ge"}
	if r := findAllConcatenatedWordsInADict(words); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
