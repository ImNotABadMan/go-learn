package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverEndpoint, err := net.Listen("tcp", ":8224")
	logCh := make(chan string)
	if err != nil {
		panic(err)
	}
	fmt.Println("Start Listen:", serverEndpoint.Addr().String())
	go logger(logCh)
	go server(serverEndpoint, logCh)
	time.Sleep(50 * time.Second)
}

func server(tcp net.Listener, ch chan string) {
	for {
		client, err := tcp.Accept()
		if err != nil {
			panic(err)
		}

		go handler(client, ch)
	}
}

func logger(log chan string) {
	for str := range log {
		fmt.Println("Log: ", str)
	}
}

func handler(client net.Conn, log chan string) {
	log <- client.RemoteAddr().String() + " connect,log log"
	fmt.Println(client.RemoteAddr().String(), " connect to server")

	_, err := client.Write([]byte(client.RemoteAddr().String() + " connect,log log"))

	client.Close()

	if err != nil {
		panic(err)
	}
}
