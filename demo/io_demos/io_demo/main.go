package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	str := strings.NewReader("test test test test")
	strW := new(strings.Builder)
	leng, _ := io.CopyN(strW, str, 10)
	fmt.Println("len", leng)

}
