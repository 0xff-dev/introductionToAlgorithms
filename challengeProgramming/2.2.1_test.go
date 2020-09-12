package challengeProgramming

import "testing"

func TestMinCoins(t *testing.T) {
	coins := []int{3, 2, 1, 3, 0, 2}
	result := MinCoins(620, coins)
	if result != 6 {
		t.Fatalf("expect 6")
	}
}
