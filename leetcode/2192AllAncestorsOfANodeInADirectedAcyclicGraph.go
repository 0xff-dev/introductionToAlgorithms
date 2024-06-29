package leetcode

func getAncestors(n int, edges [][]int) [][]int {
	adjacencyList := make([][]int, n)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		adjacencyList[to] = append(adjacencyList[to], from)
	}

	ancestorsList := [][]int{}

	for i := 0; i < n; i++ {
		ancestors := []int{}
		visited := make(map[int]bool)
		findChildren(i, adjacencyList, visited)
		for node := 0; node < n; node++ {
			if node == i {
				continue
			}
			if _, found := visited[node]; found {
				ancestors = append(ancestors, node)
			}
		}
		ancestorsList = append(ancestorsList, ancestors)
	}

	return ancestorsList
}

func findChildren(currentNode int, adjacencyList [][]int, visitedNodes map[int]bool) {
	visitedNodes[currentNode] = true

	for _, neighbour := range adjacencyList[currentNode] {
		if _, found := visitedNodes[neighbour]; !found {
			findChildren(neighbour, adjacencyList, visitedNodes)
		}
	}
}
