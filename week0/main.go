package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("1. 用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出")
	fmt.Println("https://github.com/Go-000/Go-000/issues/82")
	tcpAddr := net.TCPAddr{
		IP:   []byte{192, 168, 10, 113},
		Port: 6666,
		Zone: "",
	}

	tcp := server(&tcpAddr)
	fmt.Println(tcp.Addr().String())
	defer func() {
		tcp.Close()
		fmt.Println("tcp Close")
	}()
	var ch = make(chan string)

	gCtx := context.Background()
	gCtx = context.WithValue(gCtx, "conn", 0)
	gCtx, cancel := context.WithTimeout(gCtx, time.Second*2)
	defer cancel()

	ctx := context.Background()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	go func() {
		for {
			var errorCount int

			if errorCount > 2 {
				panic("errorCount > 2")
			}

			conn, err := tcp.AcceptTCP()
			defer func() {
				conn.Close()
				fmt.Println(conn.RemoteAddr().String() + "Conn Close")
			}()

			ctx = context.Background()
			if _, ok := ctx.Value("conn").(int); !ok {
				ctx = context.WithValue(ctx, "conn", 0)
			}
			ctx, cancel := context.WithTimeout(ctx, time.Second*10)
			ctx = context.WithValue(ctx, "conn", ctx.Value("conn").(int)+1)
			defer cancel()

			if err != nil {
				log.Println("accept error: ", err)
				errorCount++
				time.Sleep(time.Second * 1)
			}

			go func() {
				var b []byte
				fmt.Println("get connect: ", conn.RemoteAddr().String())
				if _, err := conn.Read(b); err != nil {
					log.Println("read error: ", err)
				}
				fmt.Println(string(b))
				ch <- string(b)
			}()

			go func() {
				str := <-ch
				_, err := conn.Write([]byte("Server receive succeed." + str + " Ack: OK"))
				if err != nil {
					log.Println("Send error:", err)
				}
			}()

			select {
			case <-sig:
				fmt.Println("conn close")
				ctx = context.WithValue(ctx, "conn", ctx.Value("conn").(int)-1)
				conn.Close()
			case <-ctx.Done():
				fmt.Println("conn", ctx.Value("conn").(int))
				fmt.Println(ctx.Err())
				ctx = context.WithValue(ctx, "conn", ctx.Value("conn").(int)-1)
				gCtx = context.Background()
				if ctx.Value("conn").(int) < 1 {
					conn.Close()
					gCtx, _ = context.WithTimeout(context.Background(), time.Second*10)
				}
				gCtx = context.WithValue(gCtx, "conn", ctx.Value("conn").(int))
			}
		}
	}()

	for can := true; can; {
		select {
		case <-sig:
			fmt.Println("close")
			can = false
		case <-gCtx.Done():
			//if _, ok := gCtx.Value("conn").(int); !ok {
			//	gCtx = context.WithValue(gCtx, "conn", 0)
			//	time.Sleep(time.Second * 5)
			//} else
			if i, ok := gCtx.Value("conn").(int); ok && i < 1 {
				//time.Sleep(time.Second * 5)
				//if i, ok := ctx.Value("conn").(int); ok && i < 1 {
				can = false
				//}
			}
			fmt.Println("NOW conn", gCtx.Value("conn").(int))
		}
	}

}

func server(tcpAddr *net.TCPAddr) net.TCPListener {
	tcp, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	return *tcp
}
