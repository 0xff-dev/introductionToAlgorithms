package leetcode

import (
	"reflect"
	"testing"
)

func TestWordSubsets(t *testing.T) {
	w1, w2, exp := []string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}, []string{"facebook", "google", "leetcode"}
	if r := wordSubsets(w1, w2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	w1, w2, exp = []string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}, []string{"apple", "google", "leetcode"}
	if r := wordSubsets(w1, w2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
