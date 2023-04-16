package challengeprogrammingdatastructure

// a-b b是从a开始的最远距离节点，b-c就是直径
func getDistanceArr(start, distance int, edges map[int][]int, target []int) {
	for i := 0; i < len(edges[start]); i += 2 {
		next, dis := edges[start][i], edges[start][i+1]
		if target[next] != 0 {
			continue
		}
		target[next] = distance + dis
		getDistanceArr(next, target[next], edges, target)
	}
}
func DiameterOfTree(n int, edges [][]int) int {
	target := make([]int, n)
	edgeMap := make(map[int][]int)
	for _, e := range edges {
		if _, ok := edgeMap[e[0]]; !ok {
			edgeMap[e[0]] = make([]int, 0)
		}
		edgeMap[e[0]] = append(edgeMap[e[0]], e[1], e[2])
		edgeMap[e[1]] = append(edgeMap[e[1]], e[0], e[2])
	}
	target[0] = -1
	getDistanceArr(0, 0, edgeMap, target)
	maxNode := 0
	for i := 0; i < n; i++ {
		if target[i] > target[maxNode] {
			maxNode = i
		}
		target[i] = 0
	}
	target[maxNode] = -1
	getDistanceArr(maxNode, 0, edgeMap, target)
	ans := 0
	for i := 0; i < n; i++ {
		if ans < target[i] {
			ans = target[i]
		}
	}
	return ans
}
