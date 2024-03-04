package leetcode

import "testing"

func TestBagOfTokensScore(t *testing.T) {
	tokens := []int{100}
	power := 50
	exp := 0
	if r := bagOfTokensScore(tokens, power); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	tokens = []int{200, 100}
	power = 150
	exp = 1
	if r := bagOfTokensScore(tokens, power); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	tokens = []int{100, 200, 300, 400}
	power = 200
	exp = 2
	if r := bagOfTokensScore(tokens, power); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
