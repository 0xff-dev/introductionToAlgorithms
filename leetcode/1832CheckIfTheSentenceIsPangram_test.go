package leetcode

import "testing"

func TestCheckIfPangram(t *testing.T) {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	if !checkIfPangram(sentence) {
		t.Fatalf("expect true get false")
	}

	sentence = "leetcode"
	if checkIfPangram(sentence) {
		t.Fatalf("expect false get true")
	}
}
