package leetcode

import "testing"

func TestDecodeMessage(t *testing.T) {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	exp := "this is a secret"
	if r := decodeMessage(key, message); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	key = "eljuxhpwnyrdgtqkviszcfmabo"
	message = "zwx hnfx lqantp mnoeius ycgk vcnjrdb"
	exp = "the five boxing wizards jump quickly"

	if r := decodeMessage(key, message); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
