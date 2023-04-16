package challengeprogrammingdatastructure

import "testing"

func TestCoins(t *testing.T) {
	n := 15
	coins := []int{1, 2, 7, 8, 12, 50}
	if r := Coins(coins, n); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
