package leetcode

// n <= 15
// 这么看基本搜索+cache了
func countArrangement(n int) int {
	// 第一步，不做任何cache
	used := make([]bool, n+1)
	var lookup func(int)
	ans := 0
	// 尝试将n放到某个位置
	lookup = func(index int) {
		if index > n {
			ans++
			return
		}
		// n个位置随便选择
		for idx := 1; idx <= n; idx++ {
			if !used[idx] && (index%idx == 0 || idx%index == 0) {
				used[idx] = true
				lookup(index + 1)
				used[idx] = false
			}
		}
	}
	lookup(1)
	return ans
}
