package main

import (
	"fmt"
	"goroutine_demo/channel_demo/jisuan"
	"time"
)

func main() {
	fmt.Println("channel")
	fmt.Println("发送和接收在另一端准备好之前会阻塞，这样go就不用显示给出锁和竞态变量")
	testJisuan()
}

func testJisuan() {
	fmt.Println("测试计算")

	//var i int

	ch := make(chan string, 1)
	ch <- "6"

	// 协程的队列顺序是先进后出
	fmt.Println("协程的队列顺序是先进后出")
	go jisuan.Jisuan1(ch)

	go jisuan.Jisuan2(ch)

	time.Sleep(time.Millisecond)

	fmt.Println(<-ch)

	//没数据读取(信道缓冲区为空)会抛出错误：
	fmt.Println(<-ch)
	fmt.Println("没数据读取(信道缓冲区为空)会抛出错误：使用close关闭range不会抛出错误")
	fmt.Println("循环 for v := range ch 会不断从信道接收值，直到它被关闭。")

	for v := range ch {
		fmt.Println(v)
	}
}
