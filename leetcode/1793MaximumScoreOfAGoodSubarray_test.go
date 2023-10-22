package leetcode

import "testing"

func TestMaximumScore(t *testing.T) {
    nums := []int{1,4,3,7,4,5}
    k := 3
    if r := maximumScore(nums, k); r != 15 {
        t.Fatalf("expect 15 get %d", r)
    }
    nums = []int{5,5,4,5,4,1,1,1}
    k = 0
    if r := maximumScore(nums, k); r != 20 {
        t.Fatalf("expect 20 get %d", r)
    }
}
