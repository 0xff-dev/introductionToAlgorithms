package golang

import "fmt"

// 并发素数实现
func GenNums() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func Filter(in chan int, key int) chan int {
	ch := make(chan int)
	go func() {
		for item := range in {
			if item%key != 0 {
				ch <- item
			}
		}
	}()
	return ch
}

func Prim() {
	in := GenNums()
	for c := 1; c <= 100; c++ {
		key := <-in
		fmt.Println("prim: ", key)
		in = Filter(in, key)
	}
}
