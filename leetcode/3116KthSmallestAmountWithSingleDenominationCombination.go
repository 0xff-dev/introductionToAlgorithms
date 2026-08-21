package leetcode

import "sort"

// 计算最大公约数 (Greatest Common Divisor)
func gcdf3116(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 计算最小公倍数 (Least Common Multiple)
func lcm(a, b int64) int64 {
	// 防止乘法溢出或无效值，但根据题目数据范围，int64 足够安全
	if a == 0 || b == 0 {
		return 0
	}
	return (a / gcdf3116(a, b)) * b
}

func findKthSmallest(coins []int, k int) int64 {
	// 转换成 int64 避免在大数据下溢出
	var c64 []int64
	seen := make(map[int64]bool)
	for _, c := range coins {
		if !seen[int64(c)] {
			seen[int64(c)] = true
			c64 = append(c64, int64(c))
		}
	}
	sort.Slice(c64, func(i, j int) bool {
		return c64[i] < c64[j]
	})

	n := len(c64)
	k64 := int64(k)

	// 计算在给定上限 x 内，有多少个数字可以被至少一个硬币整除（容斥原理）
	countAmounts := func(x int64) int64 {
		var total int64 = 0

		// 使用 DFS 枚举所有硬币组合的子集
		var dfs func(index int, lcmVal int64, count int)
		dfs = func(index int, lcmVal int64, count int) {
			if lcmVal > x {
				return
			}
			if index == n {
				if count > 0 {
					if count%2 == 1 {
						total += x / lcmVal
					} else {
						total -= x / lcmVal
					}
				}
				return
			}

			// 不选当前的硬币
			dfs(index+1, lcmVal, count)

			// 选择当前的硬币
			nextLcm := lcm(lcmVal, c64[index])
			if nextLcm <= x {
				dfs(index+1, nextLcm, count+1)
			}
		}

		dfs(0, 1, 0)
		return total
	}

	// 二分查找范围
	var low int64 = 1
	var high int64 = c64[0] * k64
	ans := high

	for low <= high {
		mid := low + (high-low)/2
		if countAmounts(mid) >= k64 {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}
