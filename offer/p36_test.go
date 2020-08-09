package offer

import (
	"fmt"
	"testing"
)

func TestInverseParis(t *testing.T) {
	input := []int{7, 5, 6, 3}
	fmt.Println(InverseParis(input))
	fmt.Println("result array: ", input)
}
