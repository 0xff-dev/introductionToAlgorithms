package leetcode

type UnionFind1361 struct {
	father []int
}

func (uf *UnionFind1361) FindFather(x int) int {
	if uf.father[x] != x {
		uf.father[x] = uf.FindFather(uf.father[x])
	}
	return uf.father[x]
}

// 感觉是并查集呢
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {

	father := make([]int, n)
	for i := 0; i < n; i++ {
		// 孤立的节点
		father[i] = i
	}

	for node := 0; node < n; node++ {
		left := leftChild[node]
		if left != -1 {
			lf := father[left]
			if lf != left && lf != node || left == father[node] {
				return false
			}
			father[left] = node
		}
		right := rightChild[node]
		if right != -1 {
			rf := father[right]
			if rf != right && rf != node || right == father[node] {
				return false
			}
			father[right] = node
		}
	}
	/*
			0
		1       2
		0 1 2
	*/
	sameFather := -1
	for node := 0; node < n; node++ {
		steps := 0
		walkerNode := node
		nodeFather := father[walkerNode]
		for nodeFather != walkerNode && steps <= n {
			walkerNode = nodeFather
			nodeFather = father[walkerNode]
			steps++
		}
		if steps > n {
			return false
		}
		if sameFather != -1 && nodeFather != sameFather {
			return false
		}
		sameFather = nodeFather
	}

	return true
}
