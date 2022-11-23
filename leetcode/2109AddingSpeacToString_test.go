package leetcode

import "testing"

func TestAddSpaces(t *testing.T) {
	s := "LeetcodeHelpsMeLearn"
	spaces := []int{8, 13, 15}
	if r := addSpaces(s, spaces); r != "Leetcode Helps Me Learn" {
		t.Fatalf(`expect "Leetcode Helps Me Learn" get %s`, r)
	}

	s = "icodeinpython"
	spaces = []int{1, 5, 7, 9}
	if r := addSpaces(s, spaces); r != "i code in py thon" {
		t.Fatalf(`expect "i code in py thon" get %s`, r)
	}

	s = "spacing"
	spaces = []int{0, 1, 2, 3, 4, 5, 6}
	if r := addSpaces(s, spaces); r != " s p a c i n g" {
		t.Fatalf(`expect " s p a c i n g" get %s`, r)
	}

	s = "a"
	spaces = []int{0}
	if r := addSpaces(s, spaces); r != " a" {
		t.Fatalf(`expect " a" get %s`, r)
	}
}
