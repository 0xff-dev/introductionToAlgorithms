package golang

import (
	"fmt"
	"sync"
)

var POOL = 100

func groutine1(p chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("groutine-1:", i)
		}
	}
}

func groutine2(p chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("groutine-2:", i)
		}
	}
}

func PrintNumAndChar() {
	msg := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go groutine1(msg, &wg)
	go groutine2(msg, &wg)
	wg.Wait()
}
