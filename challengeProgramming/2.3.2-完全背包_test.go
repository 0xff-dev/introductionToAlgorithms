package challengeProgramming

import (
	"fmt"
	"testing"
)

func TestCompleteBakcpack(t *testing.T) {
	weights, values := []int{3, 4, 2}, []int{4, 5, 3}
	if res := CompleteBackpack(weights, values, 7); res != 10 {
		t.Fatalf("expect 10 get %d ", res)
	}
}

func TestGeneratingFunction(t *testing.T) {
	values, nums := []int{3, 5, 8}, []int{3, 2, 2}
	for _, aim := range []int{3, 6, 9, 5, 8, 12, 17, 19} {
		fmt.Println("aim: ", aim)
		GeneratingFunction(values, nums, aim)
	}
}

type sequence struct {
	Nums []int
	Ans  int
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	for _, testCase := range []sequence{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{4, 2, 3, 1, 5}, 3},
	} {
		if res := LongestIncreasingSubsequence(testCase.Nums); res != testCase.Ans {
			t.Fatalf("expect %d get %d", testCase.Ans, res)
		}
	}
}
