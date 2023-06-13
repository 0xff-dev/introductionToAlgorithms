package leetcode

import "testing"

func TestEqualPairs(t *testing.T) {
    grid := [][]int{
        {3, 2, 1},
        {1, 7, 6},
        {2, 7, 7},
    }
    if r := equalPairs(grid); r != 1 {
        t.Fatalf("expect 1 get %d", r)
    }
    grid = [][]int{
        {3, 1, 2, 2},
        {1, 4, 4, 5},
        {2, 4, 2, 2},
        {2, 4, 2, 2},
    }
    if r := equalPairs(grid); r != 3 {
        t.Fatalf("expect 3 get %d", r)
    }
}
