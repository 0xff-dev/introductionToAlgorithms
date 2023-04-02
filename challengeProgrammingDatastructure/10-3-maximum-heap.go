package challengeprogrammingdatastructure

func maintain(index, heapSize int, nodes []int, less func(i, j int) bool) {
	root := index
	left := root*2 + 1
	right := (root + 1) * 2
	if left < heapSize && less(nodes[left], nodes[root]) {
		root = left
	}
	if right < heapSize && less(nodes[right], nodes[root]) {
		root = right
	}

	if root != index {
		nodes[root], nodes[index] = nodes[index], nodes[root]
		maintain(root, heapSize, nodes, less)
	}
}

func build(nodes []int, less func(i, j int) bool) {
	for idx := len(nodes) / 2; idx >= 0; idx-- {
		maintain(idx, len(nodes), nodes, less)
	}
}

func heapSort(nodes []int, less func(i, j int) bool) {
	build(nodes, less)
	size := len(nodes)
	for ; size > 1; size-- {
		nodes[size-1], nodes[0] = nodes[0], nodes[size-1]
		maintain(0, size-1, nodes, less)
	}
}
