package jisuan

import (
	"fmt"
)

func Jisuan1(ch chan string) string {
	i := <-ch
	i += " Jisuan 1 "
	//i := 1
	// io输出, 也是协程组装数据流
	ch <- i
	fmt.Println("jisuan 1 ")
	//time.Sleep(time.Millisecond)
	close(ch)

	return i
}

func Jisuan2(ch chan string) string {
	num := <-ch
	num += " Jisuan2"
	// 数据流的打印
	// io输出, 也是协程组装数据流
	ch <- num
	fmt.Println("jisuan 2")
	return num
}

//
//// 1. Jisuan1结束了，Jisuan2才调度起来，Jisuan2向关闭了的信道发送消息，错误
//jisuan 1
//panic: send on closed channel
//
//goroutine 7 [running]:
//goroutine_demo/channel_demo/jisuan.Jisuan2(0xc00004e180, 0x1, 0x1)
///home/ubuntu/go/code/src/goroutine_demo/channel_demo/jisuan/jisuan.go:25 +0xb4
//created by io_demo.testJisuan
///home/ubuntu/go/code/src/goroutine_demo/channel_demo/io_demo.go:27 +0x173
//exit status 2
//
//// 2. Jisuan1执行了一半，还未执行close，Jisuan2发送到信道成功
//// 结果
//jisuan 1
//jisuan 2
//6 Jisuan 1  Jisuan2
//// 3. Jisuan2执行先，执行Jisuan1，
//jisuan 2
//jisuan 1
//6 Jisuan2 Ji
