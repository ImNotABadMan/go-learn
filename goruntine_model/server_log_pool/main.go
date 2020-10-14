package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverListen, _ := net.Listen("tcp", ":8225")

	serverCh := make(chan string)

	fmt.Println("Start:", serverListen.Addr().String())

	go serverProcess(serverListen, serverCh)
	go logPool(4, serverCh)

	time.Sleep(20 * time.Second)
}

func serverProcess(s net.Listener, ch chan string) {
	for {
		conn, err := s.Accept()
		if err != nil {
			panic(err)
		}

		go handler(conn, ch)
	}
}

func logPool(loggerCount int, serverCh chan string) {
	logCh := make(chan string)
	extraCh := make(chan string)

	// 日志记录池的工作协程
	for i := 0; i < loggerCount; i++ {
		go logger(logCh, extraCh, i)
	}

	// 日志记录池中另一些额外协程
	go extraLogProcess(extraCh)

	for {
		str := <-serverCh
		logCh <- str
	}
}

func extraLogProcess(extraCh chan string) {
	for {
		str := <-extraCh
		//time.Sleep(100 * time.Millisecond)
		fmt.Println("extraLog:", str)
	}
}

func logger(ch chan string, extraCh chan string, logNo int) {
	for {
		str := <-ch
		extraCh <- str
		fmt.Println("Logger", logNo, ":", str)
		fmt.Println()
	}
}

func handler(conn net.Conn, serverCh chan string) {
	str := conn.RemoteAddr().String() + " connect to " + conn.LocalAddr().String()

	time.Sleep(100 * time.Millisecond)

	serverCh <- str
	fmt.Println(str)

	_, err := conn.Write([]byte(str))

	if err != nil {
		panic(err)
	}

	errClose := conn.Close()

	if errClose != nil {
		panic(errClose)
	}
}
