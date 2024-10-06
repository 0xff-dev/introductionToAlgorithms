package leetcode

import "testing"

func TestAreSentencesSimilar(t *testing.T) {
	sentence1, sentence2, exp := "My name is Haley", "My Haley", true
	if r := areSentencesSimilar(sentence1, sentence2); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	sentence1, sentence2, exp = "of", "A lot of words", false
	if r := areSentencesSimilar(sentence1, sentence2); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	sentence1, sentence2, exp = "Eating right now", "Eating", true
	if r := areSentencesSimilar(sentence1, sentence2); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
