package main

import (
	"bytes"
	"fmt"
	"net"
	"syscall"
	"time"
)

func main() {
	// 通信域，类型，0自动选择最好的
	socket, errS := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if errS != nil {
		fmt.Println("socket err", errS)
	}

	fmt.Println(socket)

	// 超时设置, 建立连接，什么协议
	netO, err := net.DialTimeout("tcp", "192.168.10.113:80", time.Second*2)

	if err != nil {
		fmt.Println("net err", err)
	}
	fmt.Println(netO, netO.RemoteAddr(), "<-", netO.LocalAddr())

	// 需要先请求，nginx监听的80端口才响应，net才能read导数据
	wrr, _ := netO.Write([]byte{'0'})
	if err != nil {
		fmt.Println("net Write err", err)
	}
	fmt.Println("模拟请求:", wrr)

	responseSlice := bytes.Buffer{}
	// 缓冲区，冲实现了io.Reader的接口中读取数据
	readLen, _ := responseSlice.ReadFrom(netO)

	fmt.Println(responseSlice.String())
	fmt.Println("read len:", readLen)

	fmt.Println("go run end")

	// 复用长连接，服务器没有返回新的了，直接EOF
	res, errR := netO.Read([]byte{'1'})
	fmt.Println(errR, res)

	if err := netO.Close(); err != nil {
		fmt.Println("close:", err)
	}

	fmt.Println("第二次请求：")
	netO2, errNet := net.DialTimeout("tcp", "192.168.10.113:80", time.Second*2)

	if errNet != nil {
		fmt.Println("net err", errNet)
	}

	w, errW := netO2.Write([]byte{'1', '1'})

	fmt.Println("write", w)
	if errW != nil {
		fmt.Println("write", errW)
	}

	readBytes := make([]byte, 100)
	resSlice := []byte{}

	for {
		// 读出来的数据需要持久化，Read是一个流读取
		resLen, err := netO2.Read(readBytes)
		resSlice = append(resSlice, readBytes...)

		if err != nil {
			fmt.Println("read", err)
			value, ok := err.(*net.OpError)
			fmt.Println("type", value, ok)
			if err := netO2.Close(); err != nil {
				fmt.Println("netO2 Close", err)
			}
			break
		} else {
			// 不断被初始化
			readBytes = make([]byte, resLen/2+resLen)
		}
	}

	fmt.Println("\n" + string(resSlice))
	fmt.Println("read len", len(resSlice))
	fmt.Println("read cap", cap(resSlice))

}
