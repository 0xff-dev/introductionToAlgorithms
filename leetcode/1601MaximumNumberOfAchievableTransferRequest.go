package leetcode

// 我感觉是图算法，结果 直接暴力搜索, 想暴力搜索，没敢写
func maximumRequests(n int, requests [][]int) int {
	var lookup func(int, int, []int)
	ans := 0
	lookup = func(index, count int, in []int) {
		if index == len(requests) {
			for i := 0; i < len(in); i++ {
				if in[i] != 0 {
					return
				}
			}
			if count > ans {
				ans = count
			}
			return
		}
		in[requests[index][0]]--
		in[requests[index][1]]++
		lookup(index+1, count+1, in)
		in[requests[index][0]]++
		in[requests[index][1]]--
		lookup(index+1, count, in)
	}
	lookup(0, 0, make([]int, n))
	return ans
}
