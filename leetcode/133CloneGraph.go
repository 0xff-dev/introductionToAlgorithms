package leetcode

func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	selfMap := make(map[int]*GraphNode)
	newRoot := &GraphNode{Val: node.Val, Neighbors: make([]*GraphNode, 0)}
	selfMap[node.Val] = newRoot
	visited := map[int]struct{}{}
	cloneGraphAux(node, newRoot, selfMap, visited)
	return newRoot
}

/// 记得这东西如何停止，别特么的
func cloneGraphAux(node, newRoot *GraphNode, selfMap map[int]*GraphNode, visited map[int]struct{}) {
	if _, ok := visited[node.Val]; ok {
		return
	}
	if len(node.Neighbors) > 0 {
		visited[node.Val] = struct{}{}
		for _, neighbor := range node.Neighbors {
			existNode, ok := selfMap[neighbor.Val]
			if !ok {
				existNode = &GraphNode{Val: neighbor.Val, Neighbors: make([]*GraphNode, 0)}
				selfMap[neighbor.Val] = existNode
			}
			addNode := true
			for _, n := range newRoot.Neighbors {
				if n.Val == neighbor.Val {
					addNode = false
					break
				}
			}

			if addNode {
				newRoot.Neighbors = append(newRoot.Neighbors, existNode)
				cloneGraphAux(neighbor, existNode, selfMap, visited)
			}
		}
	}

}
