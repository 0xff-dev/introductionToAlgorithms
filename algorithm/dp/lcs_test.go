package dp

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	s1, s2 := "BDCABA", "ABCBDAB"
	fmt.Println(LCS(s2, s1))
	fmt.Println("scroll array: ", ScrollArrayLCS(s2, s1))
	s1, s2 = "", "aaaaaaaa"
	fmt.Println(LCS(s2, s1))
}
