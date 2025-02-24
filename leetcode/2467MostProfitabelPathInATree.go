package leetcode

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	maxIncome := math.MinInt32
	tree := make([][]int, n)
	visited := make([]bool, n)
	bobPath := make(map[int]int)

	var nodeQueue [][3]int
	nodeQueue = append(nodeQueue, [3]int{0, 0, 0})

	// Form tree with edges
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	// Find the path taken by bob to reach node 0 and the times it takes to get there
	findBobPath(bob, 0, bobPath, visited, tree)

	// Breadth First Search
	visited = make([]bool, n)
	for len(nodeQueue) > 0 {
		sourceNode := nodeQueue[0][0]
		time := nodeQueue[0][1]
		income := nodeQueue[0][2]
		nodeQueue = nodeQueue[1:]

		// Alice reaches the node first
		if _, exists := bobPath[sourceNode]; !exists || time < bobPath[sourceNode] {
			income += amount[sourceNode]
		} else if time == bobPath[sourceNode] { // Alice and Bob reach the node at the same time
			income += amount[sourceNode] / 2
		}

		// Update max value if current node is a new leaf
		if len(tree[sourceNode]) == 1 && sourceNode != 0 {
			if income > maxIncome {
				maxIncome = income
			}
		}

		// Explore adjacent unvisited vertices
		for _, adjacentNode := range tree[sourceNode] {
			if !visited[adjacentNode] {
				nodeQueue = append(nodeQueue, [3]int{adjacentNode, time + 1, income})
			}
		}

		// Mark and remove current node
		visited[sourceNode] = true
	}

	return maxIncome
}

func findBobPath(sourceNode int, time int, bobPath map[int]int, visited []bool, tree [][]int) bool {
	// Mark and set time node is reached
	bobPath[sourceNode] = time
	visited[sourceNode] = true

	// Destination for Bob is found
	if sourceNode == 0 {
		return true
	}

	// Traverse through unvisited nodes
	for _, adjacentNode := range tree[sourceNode] {
		if !visited[adjacentNode] {
			if findBobPath(adjacentNode, time+1, bobPath, visited, tree) {
				return true
			}
		}
	}

	// If node 0 isn't reached, remove current node from path
	delete(bobPath, sourceNode)
	return false
}

