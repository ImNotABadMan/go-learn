package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 多路复用
func main() {
	//并发世界中流行的模式之一是所谓的 fan-in 模式
	ch := make(chan string)
	out := make(chan string)
	for i := 0; i < 2; i++ {
		go producer(ch, int64(i), 1)
	}
	go reader(out)

	for {
		outGlobal := <-ch
		out <- outGlobal
		strSlice := strings.Split(outGlobal, " ")
		res, _ := strconv.ParseInt(strSlice[len(strSlice)-1], 10, 10)
		if res > 10 {
			close(ch)
			close(out)
			fmt.Println("Stop")
			break
		}
	}
}

func producer(ch chan string, producerNo int64, times int) {
	out := 0
	for {
		out++
		if out > 10 && producerNo == 1 {
			break
		}
		str := "Producer " + strconv.FormatInt(producerNo, 10) + " 生产 " +
			strconv.FormatInt(int64(out), 10)

		//fmt.Println("Producer", producerNo, "put", str)
		time.Sleep(time.Duration(times+int(producerNo)) * 500 * time.Millisecond)

		ch <- str

	}

}

func reader(ch chan string) {
	// 读取通道也是 <-
	for value := range ch {
		fmt.Println("Read -- ", value)
	}
}
