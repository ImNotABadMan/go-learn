package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("协程执行，主程序退出，协程也结束")
	go say(1)
	say(2)
	time.Sleep(time.Microsecond * 100)
}

func init() {
	fmt.Println("f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中。")
	go func() int {
		res := 5 + 5 + 5
		fmt.Println("res:", res)
		return res
	}()

	fmt.Println("Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 Go 中并不经常用到")
	fmt.Println("要让协程输出，主程序需要模拟大于程序代码段执行时间io等待，不用channel的情况下")
	fmt.Println("协程是并发执行的，多CPU就有可能并行执行，所以结果是协程调度的结果，过程执行可能不一样，执行的时间片段由goroutine和系统完成")
	time.Sleep(time.Microsecond)
}

func say(no int) {
	for index := 0; index < 5; index++ {
		fmt.Println(no, "say", index)
	}
}
