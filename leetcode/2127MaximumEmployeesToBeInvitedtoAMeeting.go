package leetcode

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	inDegree := make([]int, n)

	// Calculate in-degree for each node
	for person := 0; person < n; person++ {
		inDegree[favorite[person]]++
	}

	// Topological sorting to remove non-cycle nodes
	queue := []int{}
	for person := 0; person < n; person++ {
		if inDegree[person] == 0 {
			queue = append(queue, person)
		}
	}

	depth := make([]int, n)
	// Initialize depths
	for i := range depth {
		depth[i] = 1
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		nextNode := favorite[currentNode]
		depth[nextNode] = max(depth[nextNode], depth[currentNode]+1)
		inDegree[nextNode]--
		if inDegree[nextNode] == 0 {
			queue = append(queue, nextNode)
		}
	}

	longestCycle := 0
	twoCycleInvitations := 0

	// Detect cycles
	for person := 0; person < n; person++ {
		if inDegree[person] == 0 {
			continue // Already processed
		}

		cycleLength := 0
		current := person
		for inDegree[current] != 0 {
			inDegree[current] = 0 // Mark as visited
			cycleLength++
			current = favorite[current]
		}

		if cycleLength == 2 {
			// For 2-cycles, add the depth of both nodes
			twoCycleInvitations += depth[person] + depth[favorite[person]]
		} else {
			longestCycle = max(longestCycle, cycleLength)
		}
	}

	return max(longestCycle, twoCycleInvitations)
}
