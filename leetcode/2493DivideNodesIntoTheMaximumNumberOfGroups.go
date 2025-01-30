package leetcode

func magnificentSets(n int, edges [][]int) int {
	adjList := make([][]int, n) // Create adjacency list
	parent := make([]int, n)    // Parent array for Union-Find
	depth := make([]int, n)     // Depth array for Union by Rank

	// Initialize parent array
	for i := range parent {
		parent[i] = -1
	}

	// Build the adjacency list and apply Union-Find for each edge
	for _, edge := range edges {
		adjList[edge[0]-1] = append(adjList[edge[0]-1], edge[1]-1)
		adjList[edge[1]-1] = append(adjList[edge[1]-1], edge[0]-1)
		Union(edge[0]-1, edge[1]-1, parent, depth)
	}

	numOfGroupsForComponent := make(map[int]int) // Store max groups for components

	// For each node, calculate the maximum number of groups for its component
	for node := 0; node < n; node++ {
		numberOfGroups := getNumberOfGroups(adjList, node, n)
		if numberOfGroups == -1 {
			return -1 // If invalid split, return -1
		}
		rootNode := find(node, parent)
		numOfGroupsForComponent[rootNode] = max(numOfGroupsForComponent[rootNode], numberOfGroups)
	}

	// Calculate the total number of groups across all components
	totalNumberOfGroups := 0
	for _, numberOfGroups := range numOfGroupsForComponent {
		totalNumberOfGroups += numberOfGroups
	}
	return totalNumberOfGroups
}

// Find the root of the given node in the Union-Find structure
func find(node int, parent []int) int {
	for parent[node] != -1 {
		node = parent[node]
	}
	return node
}

// Union operation to merge two sets
func Union(node1, node2 int, parent []int, depth []int) {
	node1 = find(node1, parent)
	node2 = find(node2, parent)

	// If both nodes already belong to the same set, no action needed
	if node1 == node2 {
		return
	}

	// Union by rank (depth) to keep the tree balanced
	if depth[node1] < depth[node2] {
		node1, node2 = node2, node1
	}
	parent[node2] = node1

	// If the depths are equal, increment the depth of the new root
	if depth[node1] == depth[node2] {
		depth[node1]++
	}
}

// Function to calculate the number of groups for a given component starting from srcNode
func getNumberOfGroups(adjList [][]int, srcNode, n int) int {
	nodesQueue := []int{srcNode}
	layerSeen := make([]int, n)

	// Initialize layerSeen with a default value of -1
	for i := range layerSeen {
		layerSeen[i] = -1
	}

	layerSeen[srcNode] = 0
	deepestLayer := 0

	// Perform BFS to calculate the number of layers (groups)
	for len(nodesQueue) > 0 {
		numOfNodesInLayer := len(nodesQueue)
		for i := 0; i < numOfNodesInLayer; i++ {
			currentNode := nodesQueue[0]
			nodesQueue = nodesQueue[1:]
			for _, neighbor := range adjList[currentNode] {
				// If neighbor hasn't been visited, assign it to the next layer
				if layerSeen[neighbor] == -1 {
					layerSeen[neighbor] = deepestLayer + 1
					nodesQueue = append(nodesQueue, neighbor)
				} else {
					// If the neighbor is already in the same layer, return -1 (invalid partition)
					if layerSeen[neighbor] == deepestLayer {
						return -1
					}
				}
			}
		}
		deepestLayer++
	}
	return deepestLayer
}
