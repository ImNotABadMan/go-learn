package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("start io_demo")

	// 放在panic前面，defer才会压入代码栈中，go 才会生成deter先进后出的队列执行
	defer func() {
		fmt.Println("begin defer")

		if p := recover(); p != nil {
			fmt.Println(p)
		}

		fmt.Println("end defer")
	}()

	testPanic()

	fmt.Println("end io_demo")
}

func testPanic() {

	err := errors.New("custom error to panic")

	fmt.Println("before panic")

	panic(err)

	fmt.Println("after panic")

}
