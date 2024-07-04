package leetcode

import (
	"testing"
)

func TestNumFriendRequests(t *testing.T) {
	ages := []int{16, 16}
	exp := 2
	if r := numFriendRequests(ages); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ages = []int{16, 17, 18}
	exp = 2
	if r := numFriendRequests(ages); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ages = []int{20, 30, 100, 110, 120}
	exp = 3
	if r := numFriendRequests(ages); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ages = []int{54, 23, 102, 90, 40, 74, 112, 74, 76, 21}
	exp = 22
	if r := numFriendRequests(ages); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
