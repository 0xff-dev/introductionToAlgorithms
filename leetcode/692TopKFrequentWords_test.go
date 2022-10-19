package leetcode

import (
	"reflect"
	"testing"
)

func TestTopKFrequent692(t *testing.T) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	exp := []string{"i", "love"}
	if r := topKFrequent692(words, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	words = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k = 4
	exp = []string{"the", "is", "sunny", "day"}
	if r := topKFrequent692(words, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
