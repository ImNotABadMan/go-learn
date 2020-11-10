package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	str := strings.Builder{}
	str.WriteString("test test")

	// 复制
	strCopy := str
	fmt.Println("copy", strCopy)

	str.Reset()
	strCopy1 := str
	strCopy1.Grow(8)

	fmt.Println("str reset copy1", strCopy1)

	str.WriteString("test test")

	fmt.Println("str", str)
	fmt.Println("Cap", str.Cap())
	fmt.Println("Len", str.Len())
	fmt.Println("String", str.String())

	str.Grow(8)
	fmt.Println("Grow")
	fmt.Println("str", str)
	fmt.Println("Cap", str.Cap())
	fmt.Println("Len", str.Len())
	fmt.Println("String", str.String())

	fmt.Println("\nstr.Reader")

	strR := strings.NewReader("string reader new")
	fmt.Println("str", strR)
	fmt.Println("Size", strR.Size())
	fmt.Println("Len", strR.Len())

	strReStr, size, _ := strR.ReadRune()

	fmt.Println(string(strReStr), size)
	fmt.Println("str", strR)
	fmt.Println("Size", strR.Size())
	fmt.Println("Len", strR.Len())
	fmt.Println("String", strR.UnreadRune())

	var offset int64 = 7
	seek, _ := strR.Seek(offset, io.SeekCurrent)
	expected := strR.Size() - int64(strR.Len()) + offset
	fmt.Println("io", io.SeekCurrent, "offset", offset)
	fmt.Println("seek", seek)
	fmt.Println("expected", expected)
}
