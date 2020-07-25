package dp

import (
	"fmt"
	"testing"
)

func TestMaxPrice(t *testing.T) {
	for index := 0; index <= 10; index++ {
		fmt.Println("res: ", index, " = ", MaxPrice(index))
	}
}
