package leetcode

import "testing"

func TestReplaceWords(t *testing.T) {
	dic := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	exp := "the cat was rat by the bat"
	if r := replaceWords(dic, sentence); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	dic = []string{"a", "b", "c"}
	sentence = "aadsfasf absbs bbab cadsfafs"
	exp = "a a b c"
	if r := replaceWords(dic, sentence); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
