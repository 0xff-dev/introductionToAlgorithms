package offer

import (
	"fmt"
	"testing"
)

func TestMaxSumInSubArray(t *testing.T) {
	input := []int{1, -2, 3, 10, -4, 7, 2, -5}
	res := MaxSumInSubArray(input)
	fmt.Println(res)

	input = []int{-1, -2, -3, 5}
	res = MaxSumInSubArray(input)
	fmt.Println(res)

	input = []int{-1, -2, -3, 5, -4, 6}
	res = MaxSumInSubArray(input)
	fmt.Println(res)

	input = []int{1, 2, 3, 4, 5, 6}
	res = MaxSumInSubArray(input)
	fmt.Println(res)

	input = []int{5, -6, -7, 9, 4, -5, 7, 8}
	res = MaxSumInSubArray(input)
	fmt.Println(res)
}
