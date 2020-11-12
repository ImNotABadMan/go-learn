package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	test2()
}

func test2() {
	max := 16
	sLen := 12

	sliceInt := make([]int, sLen)

	runtime.GOMAXPROCS(sLen)

	wait := sync.WaitGroup{}
	wait.Add(max)
	ch := make(chan int)

	for index := 0; index < max; index++ {
		go func(goInt []int, index int) {
			res := <-ch
			time.Sleep(time.Millisecond * 2)
			wait.Done()
			fmt.Println(index, "index", res)
		}(sliceInt, index)
	}

	// 发送信号
	for index := 0; index < max; index++ {
		ch <- index
	}

	wait.Wait()
	fmt.Println("Finish")
	fmt.Println(sliceInt)
}

func test1() {
	max := 16
	sLen := 12

	sliceInt := make([]int, sLen)

	runtime.GOMAXPROCS(sLen)

	wait := sync.WaitGroup{}
	wait.Add(max)
	//mu := sync.Mutex{}
	for index := 0; index < max; index++ {
		go func(goInt []int, index int) {
			for i := 0; i < len(sliceInt); i++ {
				if sliceInt[i] == 0 {
					//mu.Lock()

					sliceInt[i]++
					time.Sleep(time.Millisecond)
					//mu.Unlock()
					fmt.Println(index)
					break
				}
			}

			time.Sleep(time.Millisecond * 2)
			wait.Done()
		}(sliceInt, index)
	}
	wait.Wait()
	fmt.Println("Finish")
	fmt.Println(sliceInt)
}
