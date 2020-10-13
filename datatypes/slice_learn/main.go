package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("test Slice")
	s := "\u00ff\u75a5"
	fmt.Println("Slice", s)
	for index, value := range s {
		fmt.Printf("%d:%c(%s)\t", index, value, value)
	}

	fmt.Println("转换")

	// 转换，字符对应一个Unicode字符，整数表示Unicode字符
	i := []int32(s)
	fmt.Println(i)

	for index, value := range i {
		fmt.Println("index: ", index, " value :", string(value))
	}

	r := []rune(s)

	for index, value := range r {
		fmt.Println("rune index: ", index, " rune value: ", value, " string: ", string(value))
	}

	// 获取字符串长度
	length := len(s) // 字节长度
	fmt.Println("字节长度 len: ", length)

	fmt.Println("字符长度：")

	// 第一种：
	length0 := len([]int32(s))
	fmt.Println("第一种：len([]int32): ", length0)

	// 第二种：
	length1 := utf8.RuneCountInString(s)
	fmt.Println("第二种：RuneCountInString: ", length1)

	// 追加字符串
	b := []byte{1, 2, 3, 4, 5}
	appS := "你好"
	appResult := append(b, appS...)
	fmt.Println("字符串变成字节切片", appResult)

	str := "首先："
	strInt := []int32(str)
	bStr := append([]byte(str), []byte(str)...)
	strIntApp := append(strInt, []int32{22221, 65, 14681}...)
	fmt.Println("append strInt:", string(strInt))
	fmt.Println("append bStr:", string(bStr))
	fmt.Println("append strIntApp:", string(strIntApp))

	// byte 跟 int32
	int32Var := 1234
	var byteVar byte = 12

	fmt.Println("转换int<-byte：", int32(byteVar))
	fmt.Println("转换int->byte：", byte(int32Var))
}
