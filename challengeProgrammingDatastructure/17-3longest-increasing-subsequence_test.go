package challengeprogrammingdatastructure

import "testing"

func TestLongestIncreasingSubsequence(t *testing.T) {
	nums := []int{5, 1, 3, 2, 4}
	if r := LongestIncreasingSubsequence(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
