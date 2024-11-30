package leetcode

// 典型
func validArrangement(pairs [][]int) [][]int {
	adjacencyMatrix := make(map[int][]int)
	inDegree := make(map[int]int)
	outDegree := make(map[int]int)

	for _, pair := range pairs {
		start := pair[0]
		end := pair[1]
		adjacencyMatrix[start] = append(adjacencyMatrix[start], end)
		outDegree[start]++
		inDegree[end]++
	}

	result := make([]int, 0)

	var visit func(int)
	visit = func(node int) {
		for len(adjacencyMatrix[node]) > 0 {
			nextNode := adjacencyMatrix[node][0]
			adjacencyMatrix[node] = adjacencyMatrix[node][1:]
			visit(nextNode)
		}
		result = append(result, node)
	}

	startNode := -1
	for node, out := range outDegree {
		if out == inDegree[node]+1 {
			startNode = node
			break
		}
	}

	if startNode == -1 {
		startNode = pairs[0][0]
	}

	visit(startNode)

	reverse2097(result)

	pairedResult := make([][]int, 0)
	for i := 1; i < len(result); i++ {
		pairedResult = append(pairedResult, []int{result[i-1], result[i]})
	}

	return pairedResult
}

func reverse2097(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
