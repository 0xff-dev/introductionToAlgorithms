package algorith

import "testing"

func TestLRU(t *testing.T) {
	l := &lru{
		List: &pageList{
			Header: nil,
			Size:   0,
		},
		Map: map[string]*pageNode{},
	}

	l.Set(pageData(1))
	// 1
	l.List.Display()
	l.Set(pageData(2))
	// 1 2
	l.List.Display()
	l.Get("1")
	// 2 1
	l.List.Display()

	l.Set(pageData(3))
	// 2 1 3
	l.List.Display()
	l.Set(pageData(4))
	// 1 3 4
	l.List.Display()
	l.Get("3")
	// 1 4 3
	l.List.Display()

	l.Set(pageData(1))
	// 4 3 1
	l.List.Display()
}

func TestDoublyLinkList(t *testing.T) {
	l := pageList{}
	for i := 0; i < 5; i++ {
		l.Add(&pageNode{
			Pre:  nil,
			Next: nil,
			Val:  pageData(i),
		})
	}
	l.Display()
	for i := 0; i < 5; i++ {
		l.Delete(l.Header)
		l.Display()
	}
	t.Log(l.Header == nil)
}
