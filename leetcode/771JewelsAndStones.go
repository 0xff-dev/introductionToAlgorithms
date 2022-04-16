package leetcode

func numJewelsInStones(jewels string, stones string) int {
	countBytes := make(map[byte]struct{})
	for _, b := range []byte(jewels) {
		countBytes[b] = struct{}{}
	}
	cnt := 0
	for _, b := range []byte(stones) {
		if _, ok := countBytes[b]; ok {
			cnt++
		}
	}
	return cnt
}
