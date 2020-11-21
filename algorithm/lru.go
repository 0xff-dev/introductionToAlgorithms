// 最近最久未使用算法
// 1. put A --> |A|
// 2. put B --> |B A|
// 3. get A --> |A B|
// 4. put C --> |C A|
package algorith

import (
	"fmt"
	"log"
)

const maxPageSize = 3

// page content
type pageInterface interface {
	String() string
	// if need, add other functions.
	Int() int
}

type pageData int

func (p pageData) String() string {
	return fmt.Sprintf("%d", p)
}
func (p pageData) Int() int {
	return int(p)
}

// map + doubly link list
type pageNode struct {
	Pre  *pageNode
	Next *pageNode
	Val  pageInterface
}

type pageList struct {
	Header *pageNode
	Size   int
}

// Add add node to list tail.
func (pl *pageList) Add(newPage *pageNode) {

	pl.Size++
	if pl.Header == nil {
		pl.Header = newPage
		pl.Header.Next = pl.Header // circle
		pl.Header.Pre = pl.Header
		return
	}

	tail := pl.Header.Pre
	newPage.Next = pl.Header
	pl.Header.Pre = newPage
	newPage.Pre = tail
	tail.Next = newPage
}

// Delete remove a node from list.
func (pl *pageList) Delete(page *pageNode) {
	if pl.Size == 0 {
		// doubly link list has no item.
		return
	}

	isHeader := pl.Header == page

	if pl.Size == 1 {
		pl.Header, pl.Size = nil, 0
		return
	}

	pl.Size--
	page.Next.Pre = page.Pre
	page.Pre.Next = page.Next
	if isHeader {
		pl.Header = pl.Header.Next
	}
}

func (pl *pageList) Display() {
	if pl.Size == 0 {
		return
	}

	walker := pl.Header.Next
	fmt.Print(pl.Header.Val.Int())
	if pl.Size > 1 {
		for ; walker != pl.Header; walker = walker.Next {
			fmt.Print(" --> ", walker.Val.Int())
		}
	}
	fmt.Println()
}

type lru struct {
	List *pageList
	Map  map[string]*pageNode
}

func (l *lru) Get(key string) pageInterface {
	nilPage := pageData(-1)
	if v, ok := l.Map[key]; ok {
		// update order.
		l.List.Delete(v)
		l.List.Add(v)
		return v.Val
	}
	return nilPage
}

func (l *lru) Set(i pageInterface) {
	// if the result of get function isn't equal -1, the list will update
	if l.Get(i.String()).Int() == -1 {
		if l.List.Size == maxPageSize {
			log.Println("max page size")
			// max page size, delete first item and add new
			v := l.Map[l.List.Header.Val.String()]
			// delete head and add new node
			l.List.Delete(v)
		}

		newPageNode := &pageNode{
			Pre:  nil,
			Next: nil,
			Val:  i,
		}
		l.Map[i.String()] = newPageNode
		l.List.Add(newPageNode)
		log.Printf("add new key %s", i.String())
		return
	}

	log.Printf("key %s already exists", i.String())
}
