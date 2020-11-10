package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("test"))
	fmt.Println("size", reader.Size(), "buffered", reader.Buffered())
	// 缓冲区为空，拷贝到缓冲区
	fmt.Println(reader.ReadRune())
	fmt.Println("size", reader.Size(), "buffered", reader.Buffered())
	// 已经拷贝到缓冲区
	fmt.Println(reader.ReadRune())

	b, _ := reader.Peek(4)
	fmt.Println(string(b))
	fmt.Println("size", reader.Size(), "buffered", reader.Buffered())
}
