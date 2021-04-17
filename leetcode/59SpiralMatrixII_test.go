package leetcode

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	for n := 1; n <= 20; n++ {
		printR(generateMatrix(n))
	}

}

func printR(r [][]int) {
	for _, row := range r {
		fmt.Printf("%v\n", row)
	}
	fmt.Println()
}
