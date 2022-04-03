package leetcode

import "strings"

var topMap = map[int][4]string{
	1: [4]string{"M", "C", "X", "I"},
	2: [4]string{"MM", "CC", "XX", "II"},
	3: [4]string{"MMM", "CCC", "XXX", "III"},
	4: [4]string{"", "CD", "XL", "IV"},
	5: [4]string{"", "D", "L", "V"},
	6: [4]string{"", "DC", "LX", "VI"},
	7: [4]string{"", "DCC", "LXX", "VII"},
	8: [4]string{"", "DCCC", "LXXX", "VIII"},
	9: [4]string{"", "CM", "XC", "IX"},
}

func intToRoman(num int) string {
	subs := []int{1000, 100, 10, 1}
	sb := strings.Builder{}
	for idx, sub := range subs {
		quotient := num / sub
		num %= sub
		if quotient != 0 {
			sb.WriteString(topMap[quotient][idx])
		}
	}
	return sb.String()
}
