package leetcode

import (
	"testing"
)

func TestLRU(t *testing.T) {
	lru := LRUConstructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	//for w := lru.l.Front(); w != lru.l.Back(); w = w.Next() {
	//	fmt.Println("walker val: ", w.Value.(elementVal), )
	//}
	//fmt.Println("walker val: ", lru.l.Back().Value)

	t.Logf("get 1 = %d", lru.Get(1))

	lru.Put(3, 3)
	t.Logf("get 2 = %d", lru.Get(2))
	lru.Put(4, 4)
	t.Logf("get 1 = %d", lru.Get(1))
	t.Logf("get 3 = %d", lru.Get(3))
	t.Logf("get 4 = %d", lru.Get(4))
}
