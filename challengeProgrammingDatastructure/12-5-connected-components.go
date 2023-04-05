package challengeprogrammingdatastructure

import "fmt"

type uinionFind125 struct {
	father []int
}

func (u *uinionFind125) findFather(x int) int {
	f := u.father[x]
	if f != x {
		u.father[x] = u.findFather(f)
	}
	return u.father[x]
}

func (u *uinionFind125) union(x, y int) {
	x1 := u.findFather(x)
	y1 := u.findFather(y)
	if x1 < y1 {
		u.father[y] = x1
	} else {
		u.father[x] = y1
	}
}

// 这里直接用并查集解决，都是同样的道理，
// 算法里通过dfs，或者bfs讲同一个连通分量的节点标记为同样的颜色。 最后判断两个节点的颜色
// 并查集则是设置所有节点的父节点为整个区域的最小值，比较的时候直接判断父节点是否相同。
func ConnectedComponents(n int, graph [][]int, testCases [][]int) {
	u := uinionFind125{father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.father[i] = i
	}
	for _, g := range graph {
		s, t := g[0], g[1]
		fs := u.findFather(s)
		ft := u.findFather(t)
		u.union(fs, ft)
	}

	for _, tc := range testCases {
		f1 := u.findFather(tc[0])
		f2 := u.findFather(tc[1])
		if f1 == f2 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
