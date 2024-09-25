package leetcode

import (
	"reflect"
	"testing"
)

func TestSumPrefixScores(t *testing.T) {
	words, exp := []string{"abc", "ab", "bc", "b"}, []int{5, 4, 3, 2}
	if r := sumPrefixScores(words); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	words, exp = []string{"abcd"}, []int{4}
	if r := sumPrefixScores(words); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
