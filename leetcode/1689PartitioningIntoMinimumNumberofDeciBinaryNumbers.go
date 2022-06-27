package leetcode

// 11+11+10=32
func minPartitions(n string) int {
	ans := 0
	for _, b := range []byte(n) {
		if int(b-'0') > ans {
			ans = int(b - '0')
		}
	}
	return ans
}
