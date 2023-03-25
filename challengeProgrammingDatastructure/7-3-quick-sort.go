package challengeprogrammingdatastructure

import "fmt"

type card struct {
	char       byte
	num, index int
}

func (c card) String() string {
	return fmt.Sprintf("%c %d:%d", c.char, c.num, c.index)
}

func QuickSort(cards []card, left, right int) {
	if left < right {
		target := cards[left]
		pre := left
		for i := 1; i <= right; i++ {
			if cards[i].num < target.num {
				pre++
				cards[i], cards[pre] = cards[pre], cards[i]
			}
		}
		cards[left], cards[pre] = cards[pre], cards[left]
		QuickSort(cards, left, pre-1)
		QuickSort(cards, pre+1, right)
	}
}
