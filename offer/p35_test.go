package offer

import (
	"fmt"
	"testing"
)

func TestFirstNotRepeatingChar(t *testing.T) {
	str := "abaccdeff"
	fmt.Println(string(FirstNotRepeatingChar(str)))

}
