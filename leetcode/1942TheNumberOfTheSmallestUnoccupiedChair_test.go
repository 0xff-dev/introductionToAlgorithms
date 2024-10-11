package leetcode

import "testing"

func TestSmallestChair(t *testing.T) {
	times, targetFriend, exp := [][]int{{1, 4}, {2, 3}, {4, 6}}, 1, 1
	if r := smallestChair(times, targetFriend); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	times, targetFriend, exp = [][]int{{3, 10}, {1, 5}, {2, 6}}, 0, 2
	if r := smallestChair(times, targetFriend); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
