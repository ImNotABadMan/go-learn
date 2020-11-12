package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
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
