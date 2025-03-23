package leetcode

import (
	"container/heap"
	"math"
)

type Pair1976 struct {
	time int64
	node int
}

type MinHeap1976 []Pair1976

func (h MinHeap1976) Len() int           { return len(h) }
func (h MinHeap1976) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap1976) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap1976) Push(x interface{}) {
	*h = append(*h, x.(Pair1976))
}

func (h *MinHeap1976) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func countPaths1976(n int, roads [][]int) int {
	const MOD = 1e9 + 7

	// Build adjacency list
	graph := make([][]Pair1976, n)
	for _, road := range roads {
		startNode, endNode, travelTime := road[0], road[1], int64(road[2])
		graph[startNode] = append(graph[startNode], Pair1976{travelTime, endNode})
		graph[endNode] = append(graph[endNode], Pair1976{travelTime, startNode})
	}

	// Min-Heap for Dijkstra's algorithm
	minHeap := &MinHeap1976{}
	heap.Push(minHeap, Pair1976{0, 0}) // {time, node}

	// Store shortest time to each node
	shortestTime := make([]int64, n)
	for i := range shortestTime {
		shortestTime[i] = math.MaxInt64
	}
	shortestTime[0] = 0 // Distance to source is 0

	// Number of ways to reach each node in shortest time
	pathCount := make([]int, n)
	pathCount[0] = 1 // 1 way to reach node 0

	for minHeap.Len() > 0 {
		curr := heap.Pop(minHeap).(Pair1976)
		currTime, currNode := curr.time, curr.node

		// Skip outdated distances
		if currTime > shortestTime[currNode] {
			continue
		}

		for _, neighbor := range graph[currNode] {
			neighborNode, roadTime := neighbor.node, neighbor.time
			// Found a new shortest path → Update shortest time and reset path count
			if currTime+roadTime < shortestTime[neighborNode] {
				shortestTime[neighborNode] = currTime + roadTime
				pathCount[neighborNode] = pathCount[currNode]
				heap.Push(minHeap, Pair1976{shortestTime[neighborNode], neighborNode})
			} else if currTime+roadTime == shortestTime[neighborNode] {
				// Found another way with the same shortest time → Add to path count
				pathCount[neighborNode] = (pathCount[neighborNode] + pathCount[currNode]) % MOD
			}
		}
	}

	return pathCount[n-1]
}
