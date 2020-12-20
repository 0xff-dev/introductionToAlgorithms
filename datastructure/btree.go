// btree
package datastructure

import "fmt"

/*
b树特性
1. 每个节点有如下的特性
	n 是存储在节点的关键字的个数
	[n]interface 存储n个关键字
    leaf 是否是叶子节点
2. 	[n]childes 指向孩子指针
3. 节点关键字对孩子节点范围的分割
k1 <= key1 <= k2 <= key2 ...
4. 每个叶节点具有相同的深度
5. 每个节点包含的关键字有上下限，用一个被称为B树的最小度数t固定数字来表示边界, t>=2
6. 出根节点外，每个节点至少包含t-1个关键字， 最多是2t-1, 当一个节点刚好是2t-1个关键字的时候，称节点是满的。

*/

const t = 5

type bTreeNode struct {
	N       int          // key count
	Keys    []int        // key words
	Leaf    bool         // is leaf node ? true, false
	Childes []*bTreeNode // childes pointer array map[string]*bTreeNode限制了类型
}

type bTree struct {
	Root *bTreeNode
}

func (bt *bTreeNode) String() {
	fmt.Println("key:-------------")
	for i := 1; i <= bt.N; i++ {
		fmt.Print(bt.Keys[i], " ")
	}
	fmt.Println("child-------")
	for _, c := range bt.Childes {
		if c == nil {
			continue
		}
		c.String()
	}
}

func (bt *bTreeNode) merge(other *bTreeNode, k int) {
	fmt.Println("merge ----", *bt)
	fmt.Println("other node---- ", *other)
	bt.Keys[t] = k
	for idx := 1; idx <= other.N; idx++ {
		bt.Keys[t+idx] = other.Keys[idx]
		bt.Childes[t+idx-1] = other.Childes[idx-1]
	}
	bt.Childes[t+other.N] = other.Childes[other.N]
	bt.N = t+other.N
}

func (bt *bTreeNode) shiftLeft(start int) {
	for idx := start+1; idx <= bt.N; idx++ {
		bt.Keys[idx-1] = bt.Keys[idx]
		bt.Childes[idx-1] = bt.Childes[idx]
	}
	bt.Keys[bt.N], bt.Childes[bt.N] = 0, nil
	bt.N--
}

func (bt *bTreeNode) shiftRight(start int) {
	if bt.N == 2*t-1 {
		fmt.Println("full can't move")
		return
	}
	for idx := bt.N+1; idx > start; idx-- {
		bt.Keys[idx] = bt.Keys[idx-1]
		bt.Childes[idx] = bt.Childes[idx-1]
	}
	bt.Keys[start] = 0
	if start == 1 {
		bt.Childes[start-1] = nil
	}
	bt.N++
}

func newbTreeNode() *bTreeNode {
	return &bTreeNode{
		N:       0,
		Keys:    make([]int, 2*t), // max = 2*t-1, index from 1
		Leaf:    true,
		Childes: make([]*bTreeNode, 2*t), // max = 2*t, because range
	}
}

// todo insert
func (bt *bTree) Insert(k int) {
	root := bt.Root
	if root.N == 2*t-1 {
		// full
		s := newbTreeNode()
		bt.Root = s
		s.N = 0
		s.Leaf = false
		s.Childes[0] = root
		bTreeSplitChild(s, 0)
		bTreeInsertNonFull(s, k)
		return
	}
	bTreeInsertNonFull(root, k)
}


func bTreeSearch(root *bTreeNode, k int) *bTreeNode {
	if root == nil {
		return nil
	}

	index := 1
	for ; index <= root.N && k > root.Keys[index]; index++ {
	}

	if index <= root.N && k == root.Keys[index] {
		return root
	} else if root.isLeaf() {
		// leaf node
		return nil
	}

	return bTreeSearch(root.Childes[index-1], k)
}

// todo create btree, insert data into tree.
// 0 1 2 3 4 5 keys
// 0 1 2 3 4 5 child
func bTreeSplitChild(root *bTreeNode, idx int) {
	node := newbTreeNode()
	y := root.Childes[idx]
	node.Leaf = y.Leaf
	node.N = t-1

	// 分裂节点x的右侧部分放到新的节点
	for j := 1; j < t; j++ {
		node.Keys[j] = y.Keys[j+t]
		y.Keys[j+t] = 0
	}

	if !y.isLeaf() {
		// copy child
		for j := 0; j < t; j++ {
			node.Childes[j] = y.Childes[j+t]
			y.Childes[j+t] = nil
		}
	}

	y.N = t-1
	for j := root.N+1; j > idx; j-- {
		root.Childes[j] = root.Childes[j-1]
		root.Keys[j] = root.Keys[j-1]
	}
	root.Childes[idx+1] = node
	root.Keys[idx+1] = y.Keys[t]
	root.N++
	// todo write disk
}

func bTreeInsertNonFull(root *bTreeNode, k int) {
	i := root.N
	if root.isLeaf() {
		// child node
		for i >= 1 && k < root.Keys[i] {
			root.Keys[i+1] = root.Keys[i]
			i--
		}
		root.Keys[i+1] = k
		root.N++
		// todo write disk
		return
	}
	for i>=1 && k < root.Keys[i] {
		i--
	}

	if root.Childes[i].N == 2*t-1 {
		// full, split
		bTreeSplitChild(root, i)
		// Compare with the figures that have been raised
		if k > root.Keys[i] {
			i++
		}
	}
	bTreeInsertNonFull(root.Childes[i], k)
}

func (bt *bTreeNode) isLeaf() bool {
	for idx := 0; idx <= bt.N; idx++ {
		if bt.Childes[idx] != nil {
			return false
		}
	}
	return true
}

/*
1. 如果k在节点x中，切x是叶子节点，直接删除即可
2. 如果k在x节点汇总，但是x是内部节点, 做以下操作
	a. 如果x的前于k的子节点y至少包含t个关键字，找出k在以y为根的子树中的前驱k`, 递归的删除k`,在x中用k`代替k.
	b. 如果y的关键字少于t，则检查位于k后的子节点z，如果z至少包含t个关键字，找出k在z中的后继k`, 递归的删除k`, 在x中用k`代替k
	c. 否则，如果y和z的关键字个数都是t-1， 则将k和z的全部合并到y，这样x失去了k和指向z的指针，并且y现在包含2t-1个关键字，释放z，并递归的从y删除k
3. 关键字不在x节点，确定包含k的子树的根x.c[i], 如果x.c[i]只有t-1个关键字
	a. 如果他的相邻的兄弟有一个至少有t个关键字，则将x重的某个值降至到x.c[i], x.c[i]相邻的左或者右兄弟的一个关键字提上去。将相应的孩子指针移动到x.c[i]
    b. 如果相邻的兄弟都只有t-1个关键字
 */
func bTreeDeleteItem(root *bTreeNode, k int) {
	if root == nil {
		return
	}

	fmt.Println("root: ", root)
	index := 1
	for ; index <= root.N && k > root.Keys[index]; index++{}
	fmt.Println("index: ", k, "-----", index)
	if index > root.N {
		index--
	}

	if k == root.Keys[index] {
		if root.isLeaf() {
			// condition 1
			fmt.Println("here ------")
			for i := index+1; i <= root.N; i++ {
				root.Keys[i-1] = root.Keys[i]
			}
			root.N--
			return
		}

		if root.Childes[index-1].N >= t {
			// condition 2a
			// 按照节点乡下查找，最右侧的一个节点。的值
			tmpNode := nodePrecursor(root.Childes[index-1])
			fmt.Println("find precursor: ", tmpNode)
			root.Keys[index] = tmpNode.Keys[tmpNode.N]
			bTreeDeleteItem(root.Childes[index-1], tmpNode.Keys[tmpNode.N])
			return
		}

		if root.Childes[index].N >= t {
			// condition 2b
			tmpNode := nodeSuccessor(root.Childes[index])
			root.Keys[index] = tmpNode.Keys[0]
			bTreeDeleteItem(root.Childes[index], tmpNode.Keys[0])
			return
		}

		if root.Childes[index-1].N == t-1 && root.Childes[index].N == t-1 {
			// condition 2c, shift and merge
			// merge
			root.Childes[index-1].merge(root.Childes[index], k)
			// shift
			for idx := index+1; idx <= root.N; idx++ {
				root.Keys[idx-1] = root.Keys[idx]
				root.Childes[idx-1] = root.Childes[idx]
			}
			root.N--
			bTreeDeleteItem(root.Childes[index-1], k)
		}
		return
	}

	if root.isLeaf() {
		fmt.Println("root: ", root.Keys, " childs: ", root.Childes)
		fmt.Println("key ", k, " isn't in tree")
		return
	}
	tmpRoot := root.Childes[index]
	if k < root.Keys[root.N] {
		tmpRoot = root.Childes[index-1]
	}
	if tmpRoot.N == t-1 {
		// right brother
		if root.Childes[index].N >= t {
			// up
			n := root.Childes[index].Keys[1]
			children := root.Childes[index].Childes[0]
			currentVal := root.Keys[index]
			tmpRoot.Keys[tmpRoot.N+1] = currentVal
			tmpRoot.Childes[tmpRoot.N+1] = children
			tmpRoot.N++
			root.Keys[index] = n

			// shift right brother
			root.Childes[index].shiftLeft(1)
			bTreeDeleteItem(tmpRoot, k)
			return
		}
		if index > 1 {
			// when index > 1, node has left brother
			if root.Childes[index-2].N >= t {
				// up
				tn := root.Childes[index-2]
				n := tn.Keys[tn.N]
				child := tn.Childes[tn.N]
				tn.N--
				currentVal := root.Keys[index-1]
				root.Keys[index-1] = n
				tmpRoot.shiftRight(1)
				tmpRoot.Keys[0] = currentVal
				tmpRoot.Childes[0] = child
				bTreeDeleteItem(tmpRoot, k)
				return
			}
		}
		//tmpRoot.N++
		//tmpRoot.Keys[tmpRoot.N] = root.Keys[index]
		tmpRoot.merge(root.Childes[index], root.Keys[index])
		if root.N == 1 {
			root = tmpRoot
			root.Leaf = false
		}
	}

	bTreeDeleteItem(tmpRoot, k)
}

func nodePrecursor(node *bTreeNode) *bTreeNode {
	if !node.isLeaf() {
		return nodePrecursor(node.Childes[node.N])
	}
	return node
}

func nodeSuccessor(node *bTreeNode) *bTreeNode {
	if !node.isLeaf() {
		return nodeSuccessor(node.Childes[0])
	}
	return node
}
