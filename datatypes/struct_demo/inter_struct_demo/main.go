package main

import (
	"datatypes/struct_demo/inter_struct_demo/my_inter_struct"
	"errors"
	"fmt"
)

func main() {
	fmt.Println("错误处理")
	err := errors.New("error string")

	fmt.Println(err)

	var errorC error

	myErr := new(my_inter_struct.MyError)

	// 方法接收值为结构体值类型，方法的接收者为副本
	// 所有给定类型的方法都应该有值或指针接收者
	(myErr).SetPointerPrefix("pointer test")
	(myErr).SetStructPrefix("struct test")

	fmt.Println((myErr).Error())
	// 接口方法是结构体实现方法的带指针的接收者 就要把结构体的指针赋值给接口类型，接口才能正确执行方法
	errorC = *myErr
	fmt.Println(errorC)
}
