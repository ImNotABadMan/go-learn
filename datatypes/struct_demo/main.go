package main

import "fmt"

func main()  {
	// 定义结构体
	type demo struct {
		id int64
		name string
	}

	// test1 是指针，指向struct类型
	test1 := new(demo)

	test1.id = 1
	test1.name = "test1"

	fmt.Println(test1)

	var test2 *demo
	test2 = new(demo)
	test2.id = 2
	test2.name = "test2"

	fmt.Println(test2)

	// 值类型
	var test3 demo
	test3.name = "test3"
	fmt.Println(test3)

	test4 := demo{
		id:   4,
		name: "test",
	}
	fmt.Println(test4)
}
