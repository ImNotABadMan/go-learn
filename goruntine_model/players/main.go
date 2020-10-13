package main

import (
	"fmt"
	"time"
)

func main() {
	ball := make(chan int)
	fmt.Println("Player Start")

	for i := 1; i < 2; i++ {
		go player(ball, 1, i)
	}

	ball <- 0
	time.Sleep(10 * time.Second)
	<-ball
}

func player(ch chan int, timeDuration int64, playerNo int) {
	for {
		times := <-ch
		times++
		if times > 21 {
			break
		}
		fmt.Println("Player ", playerNo, " kit ball ", times)
		time.Sleep(time.Duration(timeDuration) * 500 * time.Millisecond)
		ch <- times
	}
}
