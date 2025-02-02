package leetcode

import (
	"fmt"
	"sync"
	"testing"
)

func af() {
	fmt.Print("foo")
}

func bf() {
	fmt.Print("bar")
}

func TestFooBar(t *testing.T) {
	for i := 1; i < 5; i++ {
		a := NewFooBar(i)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			a.Foo(af)
		}()
		go func() {
			defer wg.Done()
			a.Bar(bf)
		}()
		wg.Wait()
		fmt.Println()
	}
}
