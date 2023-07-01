package leetcode

import "testing"

func TestDistributeCookies(t *testing.T) {
	cookies := []int{8, 15, 10, 20, 8}
	k := 2
	if r := distributeCookies(cookies, k); r != 31 {
		t.Fatalf("expect 31 get %d", r)
	}

	cookies = []int{6, 1, 3, 2, 2, 4, 1, 2}
	k = 3
	if r := distributeCookies(cookies, k); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
