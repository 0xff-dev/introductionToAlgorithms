package leetcode

import "fmt"

func commonChars(words []string) []string {
	l := len(words)
	count := make([][26]int, len(words))
	for i, word := range words {
		for _, b := range word {
			count[i][b-'a']++
		}
	}

	ans := make([]string, 0)
	for i := 0; i < 26; i++ {
		c := -1
		for j := 0; j < l; j++ {
			if count[j][i] == 0 {
				c = -1
				break
			}
			if c == -1 || c > count[j][i] {
				c = count[j][i]
			}
		}
		if c != -1 {
			for k := 0; k < c; k++ {
				ans = append(ans, fmt.Sprintf("%c", byte(i)+'a'))
			}
		}
	}
	return ans
}
