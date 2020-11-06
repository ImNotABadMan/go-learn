package main

import (
	"fmt"
	"sync"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(5)
	for index := 0; index < 4; index++ {
		go genStr(group.Done)
	}

	once := sync.Once{}
	go once.Do(func() {
		fmt.Println("sync.do once")
		group.Done()
	})
	// 	group.Add(6)发生死锁，因为once.Do中的函数只会执行一次，所以6个协程，发生5次group.Done
	go once.Do(func() {
		fmt.Println("sync.do once")
		group.Done()
	})

	group.Wait()
}

func genStr(done func()) {
	fmt.Println("str")
	fmt.Println(done)
	done()

}
