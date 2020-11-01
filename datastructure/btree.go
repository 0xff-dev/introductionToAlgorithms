// btree
package datastructure

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
type bTreeNode struct {
	N       int          // key count
	Keys    []int        // key words
	Leaf    bool         // is leaf node ? true, false
	Childes []*bTreeNode // childes pointer array map[string]*bTreeNode限制了类型
}

func bTreeSearch(root *bTreeNode, k int) *bTreeNode {
	if root == nil {
		return nil
	}

	index := 0
	for ; index < root.N && k < root.Keys[index]; index++ {
	}

	if index < root.N && k == root.Keys[index] {
		return root
	} else if root.Leaf {
		// leaf node
		return nil
	}

	return bTreeSearch(root.Childes[index], k)
}

// todo create btree, insert data into tree.
