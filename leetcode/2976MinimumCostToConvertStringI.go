package leetcode

const unreachable = 1 << 32

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	// 图算法
	// original 和 changed 组成一个单向图,可能有环图
	// 感觉是多次利用dijkstra，但是我觉得一个问题：会超时

	cache := make([][]int64, 26)
	for i := 0; i < 26; i++ {
		cache[i] = make([]int64, 26)
		for j := 0; j < 26; j++ {
			cache[i][j] = unreachable
		}
		cache[i][i] = 0
	}
	// 定义图
	for i := 0; i < len(original); i++ {
		x, y := original[i]-'a', changed[i]-'a'
		cache[x][y] = min(cache[x][y], int64(cost[i]))
	}

	// Floyd-Warshall
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			if cache[i][k] == unreachable {
				continue
			}
			for j := 0; j < 26; j++ {
				if cache[k][j] == unreachable {
					continue
				}
				cache[i][j] = min(cache[i][j], cache[i][k]+cache[k][j])
			}
		}
	}
	ans := int64(0)
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			r := cache[source[i]-'a'][target[i]-'a']
			if r == unreachable {
				return -1
			}
			ans += r
		}
	}
	return ans
}
