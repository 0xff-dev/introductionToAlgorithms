package leetcode

import "testing"

func TestFindJudge(t *testing.T) {
	n, trust := 2, [][]int{{1, 2}}
	if r := findJudge(n, trust); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	n, trust = 3, [][]int{{1, 3}, {2, 3}}
	if r := findJudge(n, trust); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	n, trust = 3, [][]int{{1, 3}, {2, 3}, {3, 1}}
	if r := findJudge(n, trust); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
