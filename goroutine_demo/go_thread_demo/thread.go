package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(new(sync.Mutex))
	var num int = 1
	go thread1(cond, &num)
	go thread2(cond, &num)
	time.Sleep(time.Millisecond * 10)
}

func thread1(c *sync.Cond, num *int) {
	c.L.Lock()
	*num++
	fmt.Println("thread1 number is ", *num)
	c.L.Unlock()
	// 需要让别的goroutine先起来，happen before
	time.Sleep(time.Millisecond)
	// 唤醒goroutine
	c.Signal()
}

func thread2(c *sync.Cond, num *int) {
	c.L.Lock()
	c.Wait()
	*num++
	fmt.Println("thread2 number is ", *num)
	c.L.Unlock()
}
