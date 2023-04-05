package challengeprogrammingdatastructure

func SingleSourceShortestPath(n int, table []adjTable) ([]int, []int) {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = INF
		}
	}
	totalEdges := 0
	for _, item := range table {
		totalEdges += item.k
		for i := 0; i < len(item.points); i += 2 {
			matrix[item.u][item.points[i]] = item.points[i+1]
		}
	}
	ans1 := BellmanFord(matrix, n, totalEdges)
	ans2 := Dijkstra(matrix, n)
	return ans1, ans2
}

func bellmanFordRelax(distance []int, from, to, weight int) {
	if distance[to] == INF {
		distance[to] = weight
		if distance[from] != INF {
			distance[to] += distance[from]
		}
		return
	}
	if distance[from] != INF && distance[from]+weight < distance[to] {
		distance[to] = distance[from] + weight
	}
}

func BellmanFord(matrix [][]int, n, totalEdges int) []int {
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = INF
	}
	distance[0] = 0
	for to := 1; to < n; to++ {
		if matrix[0][to] == INF {
			continue
		}
		distance[to] = matrix[0][to]
	}
	for i := 0; i < totalEdges-1; i++ {
		for from := 0; from < n; from++ {
			for to := 0; to < n; to++ {
				if matrix[from][to] == INF {
					continue
				}
				bellmanFordRelax(distance, from, to, matrix[from][to])
			}
		}
	}

	return distance
}

// dijkstra 算法运行中维持的关键信息是一组节点集合S
// 从起点到集合中的每个点之间的最短路径已经找到
func Dijkstra(matrix [][]int, n int) []int {
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = INF
	}
	distance[0] = 0
	visited := make([]bool, n)
	for i := 0; i < n-1; i++ {
		// 这个部分可以优化成heap
		selectNode := -1
		for from, dis := range distance {
			if dis == INF {
				continue
			}
			if !visited[from] && (selectNode == -1 || distance[from] < distance[selectNode]) {
				selectNode = from
			}
		}
		visited[selectNode] = true

		for next, dis := range matrix[selectNode] {
			if dis == INF {
				continue
			}
			if distance[next] == INF || dis+distance[selectNode] < distance[next] {
				distance[next] = dis + distance[selectNode]
			}
		}
	}
	return distance
}
