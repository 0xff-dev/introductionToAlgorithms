package challengeprogrammingdatastructure

const INF = 0xffff

type taskDesc struct {
	name  string
	times int
}

type DLNode struct {
	pre, next *DLNode
	val       int
}

// 实现一个方法，支持插入字符串，查询字符串
// 字符串仅包含 ACGT4种字符。
// 这可以考虑散列，前缀树
type op struct {
	opName string
	target string
}

type point struct {
	x, y float32
}

type card struct {
	char       byte
	num, index int
}

type treeDefine struct {
	id    int
	k     int
	child []int
}

type tree83 struct {
	id          int
	left, right int
}

type tree84 struct {
	id, left, right int
}

type BSTNode struct {
	Val         int
	Left, Right *BSTNode
}

type Matrix struct {
	row, col int
}

type adjTable struct {
	u, k   int
	points []int
}

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
