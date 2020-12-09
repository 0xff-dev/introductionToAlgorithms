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
	} else if root.Leaf {
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
	}

	if !y.Leaf {
		// copy child
		for j := 0; j < t; j++ {
			node.Childes[j] = y.Childes[j+t]
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
	if root.Leaf {
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
