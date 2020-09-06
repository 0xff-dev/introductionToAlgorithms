/*
	由若干木棍，每个长度为a[i],选择出能组成三角形的最长周长，如果不存在返回0
*/
package challengeProgramming

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TrianglePerimeter(woods []int) int {
	maxPerimeter := 0
	for i := 0; i < len(woods); i++ {
		for j := i + 1; j < len(woods); j++ {
			for u := j + 1; u < len(woods); u++ {
				// perimeter
				rest := woods[i] + woods[j] + woods[u]
				maxWood := max(woods[i], max(woods[j], woods[u]))
				// 两边之和大于第三边
				if rest-maxWood > maxWood {
					maxPerimeter = max(maxPerimeter, rest)
				}
			}
		}
	}
	return maxPerimeter
}
