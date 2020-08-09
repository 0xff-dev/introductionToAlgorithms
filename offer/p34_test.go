package offer

import (
	"fmt"
	"testing"
)

func TestUglyNum(t *testing.T) {
	for _, n := range []int{2, 3, 5, 6, 1500} {
		fmt.Println("ugly: ", UglyNum(n))
	}
}
