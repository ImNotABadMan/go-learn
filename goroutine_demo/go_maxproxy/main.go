package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	testMaxProxy()
	go longWait()
	go ShortWait()
	mainWait()
}

func testMaxProxy() {
	// 接受命令行参数
	var coreNum = flag.Int("core", 1, "cores number")
	flag.Parse()

	runtime.GOMAXPROCS(*coreNum)
}

func longWait() {
	fmt.Println("long wait start")
	time.Sleep(time.Second * 2)
	fmt.Println("long wait end")
}

func ShortWait() {
	fmt.Println("short wait start")
	time.Sleep(time.Second * 2)
	fmt.Println("short wait end")
}

func mainWait() {
	fmt.Println("io_demo wait start")
	time.Sleep(time.Second * 4)
	fmt.Println("io_demo wait end")
}
