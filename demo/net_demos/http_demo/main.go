package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
)

type ServerHandler struct{}

func (handler *ServerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Header)
	if wLen, err := writer.Write([]byte("this is go 语言")); err != nil {
		fmt.Println(err, wLen)
	}
}

func main() {
	content, _ := http.Get("http://127.0.0.1")
	buf := bytes.Buffer{}

	bodyLen, _ := buf.ReadFrom(content.Body)

	fmt.Println("body len", bodyLen)
	fmt.Println(len(buf.String()))
	fmt.Println(content.Header)

	listen, errListen := net.Listen("tcp", ":9012")
	if errListen != nil {
		fmt.Println(errListen)
	}

	err := http.Serve(listen, new(ServerHandler))
	if err != nil {
		fmt.Println("serve:", err)
	}
}
