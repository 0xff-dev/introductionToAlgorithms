package leetcode

import "testing"

func TestPalindromePartition(t *testing.T){
    s, k := "abc", 2
    if r := palindromePartition(s, k); r != 1 {
        t.Fatalf("expect 1 get %d", r)
    }

    s, k = "aabbc", 3
    if r := palindromePartition(s, k); r != 0 {
        t.Fatalf("expect 0 get %d", r)
    }

    s,k = "leetcode", 8
    if r := palindromePartition(s, k); r != 0 {
        t.Fatalf("expect 0 get %d", r)
    }
}
