package main

import (
	"fmt"
	"second/learn"
)

func init() {
	fmt.Println('\x01', "init tests")
}

func main() {
	// 调用包的 内容
	learn.Learn()
	learn.Hello()

	str := "test String"
	for i, value := range str {
		fmt.Println("now is : ", i, ", value is : ", string(value))
	}
	var test uint8
	test = 255
	fmt.Println("Int8 : ", test, " To Int16: ", int16(test), " To byte", byte(test),
		" To int8: ", int8(test))

	var test2 uint16
	test2 = 65535
	fmt.Println("Int 16: ", test2,
		" To Int32: ", int32(test2), " To init8: ", int8(test2))
}
