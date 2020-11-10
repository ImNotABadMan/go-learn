package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	fmt.Println("并发抢变量")
	sliceOnce := make([]sync.Once, 10)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)

	countCh := make(chan int)
	lock := sync.Mutex{}

	for count, index := 1, 0; index < 1000; index++ {
		go func(i int) {
			if num := rand.Intn(100); num < 10 {
				sliceOnce[num].Do(func() {
					// 并发打印出来，有延迟
					lock.Lock()
					fmt.Println("User", i, "Get Gift ", count)
					// 并发的时候，这里可能还没有执行到
					count++
					lock.Unlock()
					countCh <- count
					waitGroup.Done()
				})
			}
		}(index)

	}

	go func() {
		// 有协程进行接收通道，通道可以不需要缓存，没有协程可以立刻接收通道的，通道需要缓冲区，防止阻塞
		for index := 0; index < 10; index++ {
			fmt.Println("Gift", <-countCh)
		}
	}()

	waitGroup.Wait()
	// 这里通道需要缓冲区
	//countCh := make(chan int, 10)
	//for index := 0; index < 10; index++ {
	//	fmt.Println("Gift", <- countCh)
	//}
}
