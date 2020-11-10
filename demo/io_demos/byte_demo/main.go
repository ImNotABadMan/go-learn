package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("字节缓冲区")
	buffer := bytes.Buffer{}
	fmt.Println(buffer)

	buffer.WriteString("buffer string")
	fmt.Println("str", buffer, buffer.String())
	fmt.Println("len", buffer.Len(), "cap", buffer.Cap())

	b := make([]byte, 8)
	bR, _ := buffer.Read(b)
	fmt.Println(buffer.UnreadByte())

	fmt.Println("bR", bR)
	fmt.Println("str", buffer, buffer.String())
	fmt.Println("len(未读的长度)", buffer.Len(), "cap", buffer.Cap())

	buffer.WriteString("buffer string")
	fmt.Println("str", buffer, buffer.String())
	fmt.Println("len", buffer.Len(), "cap", buffer.Cap())

}
