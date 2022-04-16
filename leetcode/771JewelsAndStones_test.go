package leetcode

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	j, s := "aA", "aAAbbbb"
	if r := numJewelsInStones(j, s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
