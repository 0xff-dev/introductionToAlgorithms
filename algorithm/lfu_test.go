package algorith

import (
	"testing"
	"time"
)

func TestLfuList(t *testing.T) {
	l := NewLFUList()
	n1 := &lfuNode{
		Fre:     1,
		Val:     2,
		AddTime: time.Now(),
		Next:    nil,
	}
	n2 := &lfuNode{
		Fre:     2,
		Val:     3,
		AddTime: time.Now(),
		Next:    nil,
	}
	n3 := &lfuNode{
		Fre:     18,
		Val:     20,
		AddTime: time.Now(),
		Next:    nil,
	}
	n4 := &lfuNode{
		Fre:     18,
		Val:     1111,
		AddTime: time.Now(),
		Next:    nil,
	}
	l.Add(n1)
	l.Add(n2)
	l.Add(n3)
	l.Add(n4)
	l.Dis()
	l.Delete(n4)
	l.Dis()
	t.Log("delete node n2")
	l.Delete(n2)
	l.Dis()
	t.Log("delete node n1")
	l.Delete(n1)
	l.Dis()
}