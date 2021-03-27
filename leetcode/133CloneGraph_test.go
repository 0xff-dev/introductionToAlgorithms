package leetcode

import (
	"fmt"
	"testing"
)

func graphDis(node *GraphNode, visited map[int]struct{}) {
	if _, ok := visited[node.Val]; ok {
		return
	}
	visited[node.Val] = struct{}{}
	if len(node.Neighbors) == 0 {
		fmt.Printf("%d don't have neighobrs.", node.Val)
		return
	}
	for _, nei := range node.Neighbors {
		fmt.Printf("%d-----%d\n", node.Val, nei.Val)
		graphDis(nei, visited)
	}
}

func TestCloneGraph(t *testing.T) {
	graph := &GraphNode{1, make([]*GraphNode, 0)}
	graph1 := &GraphNode{2, make([]*GraphNode, 0)}
	graph2 := &GraphNode{3, make([]*GraphNode, 0)}
	graph3 := &GraphNode{4, make([]*GraphNode, 0)}
	graph.Neighbors = []*GraphNode{graph1, graph3}
	graph1.Neighbors = []*GraphNode{graph, graph2}
	graph2.Neighbors = []*GraphNode{graph1, graph3}
	graph3.Neighbors = []*GraphNode{graph, graph2}

	vis := make(map[int]struct{})
	graphDis(graph, vis)

	t.Log("clone test...\n")
	clone := cloneGraph(graph)
	visited := make(map[int]struct{})
	graphDis(clone, visited)

	nGraph := &GraphNode{Val: 1, Neighbors: nil}
	clone = cloneGraph(nGraph)
	nVis := make(map[int]struct{})
	graphDis(clone, nVis)
}
