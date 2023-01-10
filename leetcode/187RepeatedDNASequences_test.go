package leetcode

import (
	"reflect"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	exp := []string{"AAAAACCCCC", "CCCCCAAAAA"}
	if r := findRepeatedDnaSequences(s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	s = "AAAAAAAAAAAAA"
	exp = []string{"AAAAAAAAAA"}
	if r := findRepeatedDnaSequences(s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
