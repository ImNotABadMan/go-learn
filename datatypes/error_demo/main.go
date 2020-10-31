package main

import (
	"datatypes/struct_demo/inter_struct_demo/my_inter_struct"
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	testError()
}

func testError() {
	var err error

	errNew := errors.New("test error")

	fmt.Println(errNew)
	fmt.Println(err)

	err = errNew

	fmt.Println(err)

	if err != nil {
		fmt.Println(os.PathError{"Op", "D:", err})
	}

	//test
	switch t := err.(type) {
	case net.Error:
		fmt.Println(t)
		// 怎样实现接口，接口方法的接收者是什么类型的，指针还是值类型
	case *os.LinkError:
		fmt.Println(t, "os")
	case my_inter_struct.MyError:
		fmt.Println("int")
	default:
		fmt.Println("default", t)
	}
}
