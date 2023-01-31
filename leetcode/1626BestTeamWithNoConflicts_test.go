package leetcode

import "testing"

func TestBestTeamScore(t *testing.T) {
	scores, ages := []int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}
	if r := bestTeamScore(scores, ages); r != 34 {
		t.Fatalf("expect 34 get %d", r)
	}

	scores, ages = []int{4, 5, 6, 5}, []int{2, 1, 2, 1}
	if r := bestTeamScore(scores, ages); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}

	scores, ages = []int{319776, 611683, 835240, 602298, 430007, 574, 142444, 858606, 734364, 896074}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if r := bestTeamScore(scores, ages); r != 5431066 {
		t.Fatalf("expect 5431066 get %d", r)
	}

	scores, ages = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{811, 364, 124, 873, 790, 656, 581, 446, 885, 134}
	if r := bestTeamScore(scores, ages); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	scores, ages = []int{1, 3, 7, 3, 2, 4, 10, 7, 5}, []int{4, 5, 2, 1, 1, 2, 4, 1, 4}
	if r := bestTeamScore(scores, ages); r != 29 {
		t.Fatalf("expect 29 get %d", r)
	}
}
