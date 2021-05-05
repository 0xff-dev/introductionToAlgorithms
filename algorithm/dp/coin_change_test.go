package dp

import (
    "testing"
)

func TestMinCoinChange(t *testing.T) {
    target := 27
    if r := minCoins(target); r != 5 {
        t.Fatalf("expect 5 get %d", r)
    }
}
