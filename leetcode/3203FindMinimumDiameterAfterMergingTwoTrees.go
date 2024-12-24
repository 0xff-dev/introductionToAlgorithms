package leetcode

import "math"

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	// Calculate the number of nodes for each tree
	n := len(edges1) + 1
	m := len(edges2) + 1

	// Build adjacency lists for both trees
	adjList1 := buildAdjList(n, edges1)
	adjList2 := buildAdjList(m, edges2)

	// Calculate the diameters of both trees
	diameter1 := findDiameter(n, adjList1)
	diameter2 := findDiameter(m, adjList2)

	// Calculate the longest path that spans across both trees
	combinedDiameter := int(math.Ceil(float64(diameter1)/2.0)) + int(math.Ceil(float64(diameter2)/2.0)) + 1

	// Return the maximum of the three possibilities
	return max(diameter1, max(diameter2, combinedDiameter))
}

func buildAdjList(size int, edges [][]int) [][]int {
	adjList := make([][]int, size)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	return adjList
}

// Function to find the diameter of a tree using two BFS calls
// 需要学习这一个算法
func findDiameter(n int, adjList [][]int) int {
	farthestNode, _ := findFarthestNode(n, adjList, 0)

	_, diameter := findFarthestNode(n, adjList, farthestNode)
	return diameter
}

// BFS helper function to find the farthest node and its distance from the source
func findFarthestNode(n int, adjList [][]int, sourceNode int) (int, int) {
	nodesQueue := []int{sourceNode}
	visited := make([]bool, n)
	visited[sourceNode] = true

	maximumDistance := 0
	farthestNode := sourceNode

	// Explore neighbors
	for len(nodesQueue) > 0 {
		size := len(nodesQueue)
		for i := 0; i < size; i++ {
			currentNode := nodesQueue[0]
			nodesQueue = nodesQueue[1:]
			farthestNode = currentNode

			for _, neighbor := range adjList[currentNode] {
				if !visited[neighbor] {
					visited[neighbor] = true
					nodesQueue = append(nodesQueue, neighbor)
				}
			}
		}
		if len(nodesQueue) > 0 {
			maximumDistance++
		}
	}
	return farthestNode, maximumDistance
}
