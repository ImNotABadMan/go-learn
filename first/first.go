package main

import (
	"fmt"
)
import "first-moudle"
import "log"

var test int32

func main() {
	test = 10
	inTest := 20
	fmt.Println("test" + string(test))
	fmt.Println(inTest)
	var message = first_moudle.Say("my test")
	fmt.Println(message)

	// 设置Log

	log.SetPrefix("first_moudle: ")
	log.SetFlags(0)

	_, err := first_moudle.SetError("")
	if err != nil {
		//log.Fatal(err)
	}

	fmt.Println(first_moudle.SetError("test error not empty"))

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 切片
	slice := []int32{1, 2, 3}
	// 数组
	for index, value := range slice {
		fmt.Printf("%d:%d\t", index, value)
	}

	fmt.Println()
	// 数组
	arr := [...]int32{1, 2, 3, 4}

	// 数组
	for index, value := range arr {
		fmt.Printf("%d:%d\t", index, value)
	}

	stu := first_moudle.StuSrtuct{
		Id:   0,
		Name: "Test",
	}

	stu.Id = 10

	fmt.Println(stu)

	first_moudle.StructShow()

	first_moudle.ArrShow()

	first_moudle.Gs = "st"

	fmt.Println(first_moudle.Gs)

	first_moudle.GetSlice()

	_, _ = first_moudle.SetSlice([]string{"1", "2", "3", "4"})

	first_moudle.GetSlice()
}
