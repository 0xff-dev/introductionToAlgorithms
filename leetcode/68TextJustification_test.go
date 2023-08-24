package leetcode

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	exp := []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}
	if r := fullJustify(words, maxWidth); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth = 20
	exp = []string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}
	if r := fullJustify(words, maxWidth); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	words = []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	maxWidth = 16
	exp = []string{
		"ask   not   what", "your country can", "do  for  you ask", "what  you can do", "for your country",
	}
	if r := fullJustify(words, maxWidth); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
