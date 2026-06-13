package leetcode

import "bytes"

func mapWordWeights(words []string, weights []int) string {
	var (
		buf bytes.Buffer
		sum int
	)
	for i := range words {
		sum = 0
		for b := range words[i] {
			sum += weights[words[i][b]-'a']
		}
		buf.WriteByte('a' + byte(25-sum%26))
	}
	return buf.String()
}
