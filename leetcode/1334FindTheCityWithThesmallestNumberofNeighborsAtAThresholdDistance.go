package leetcode

import (
	"container/heap"
	"math"
)

// Edge represents a graph edge
type Edge struct {
	node   int
	weight int
}

// MinHeap is a min-heap of pairs (distance, node)
type MinHeap [][2]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// findTheCity finds the city with the fewest reachable cities within the distance threshold
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Adjacency list to store the graph
	adjacencyList := make([][]Edge, n)
	// Matrix to store shortest path distances from each city
	shortestPathMatrix := make([][]int, n)

	// Initialize adjacency list and shortest path matrix
	for i := 0; i < n; i++ {
		shortestPathMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				shortestPathMatrix[i][j] = math.MaxInt32
			}
		}
	}

	// Populate the adjacency list with edges
	for _, edge := range edges {
		start := edge[0]
		end := edge[1]
		weight := edge[2]
		adjacencyList[start] = append(adjacencyList[start], Edge{end, weight})
		adjacencyList[end] = append(adjacencyList[end], Edge{start, weight}) // For undirected graph
	}

	// Compute shortest paths from each city using Dijkstra's algorithm
	for i := 0; i < n; i++ {
		dijkstra(n, adjacencyList, shortestPathMatrix[i], i)
	}

	// Find the city with the fewest number of reachable cities within the distance threshold
	return getCityWithFewestReachable(n, shortestPathMatrix, distanceThreshold)
}

// dijkstra's algorithm to find shortest paths from a source city
func dijkstra(n int, adjacencyList [][]Edge, shortestPathDistances []int, source int) {
	// Priority queue to process nodes with the smallest distance first
	priorityQueue := &MinHeap{}
	heap.Init(priorityQueue)
	heap.Push(priorityQueue, [2]int{0, source})
	for i := range shortestPathDistances {
		shortestPathDistances[i] = math.MaxInt32
	}
	shortestPathDistances[source] = 0 // Distance to source itself is zero

	// Process nodes in priority order
	for priorityQueue.Len() > 0 {
		top := heap.Pop(priorityQueue).([2]int)
		currentDistance := top[0]
		currentCity := top[1]
		if currentDistance > shortestPathDistances[currentCity] {
			continue
		}

		// Update distances to neighboring cities
		for _, neighbor := range adjacencyList[currentCity] {
			neighborCity := neighbor.node
			edgeWeight := neighbor.weight
			if shortestPathDistances[neighborCity] > currentDistance+edgeWeight {
				shortestPathDistances[neighborCity] = currentDistance + edgeWeight
				heap.Push(priorityQueue, [2]int{shortestPathDistances[neighborCity], neighborCity})
			}
		}
	}
}

// getCityWithFewestReachable determines the city with the fewest number of reachable cities within the distance threshold
func getCityWithFewestReachable(n int, shortestPathMatrix [][]int, distanceThreshold int) int {
	cityWithFewestReachable := -1
	fewestReachableCount := n

	// Count number of cities reachable within the distance threshold for each city
	for i := 0; i < n; i++ {
		reachableCount := 0
		for j := 0; j < n; j++ {
			if i != j && shortestPathMatrix[i][j] <= distanceThreshold {
				reachableCount++
			}
		}
		// Update the city with the fewest reachable cities
		if reachableCount <= fewestReachableCount {
			fewestReachableCount = reachableCount
			cityWithFewestReachable = i
		}
	}
	return cityWithFewestReachable
}
