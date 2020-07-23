package offer

import (
	"fmt"
	"testing"
)

func TestReorderOddEven(t *testing.T) {
	input := []int{1, 4, 8, 2, 7, 9, 3, 6}
	ReorderOddEven(input, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(input)
}
