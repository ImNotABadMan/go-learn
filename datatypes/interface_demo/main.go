package main

import (
	"datatypes/interface_demo/test_inter"
	"fmt"
)

func main() {
	testInterface()
	testNil()
	testAnyInterface()
	testType(2)
	testType("str")
	testString()
	testTag()
}

func testInterface() {
	base := test_inter.BaseStruct{}
	base.Show()

	var baseInterface test_inter.BaseInter

	// testInter中，因为接收者为指针，接口只为（指针类型）定义，因此（值类型）并未实现 接口。
	baseInterface = &base
	fmt.Println(baseInterface.Show())

	baseInt := test_inter.MyInt(2)
	baseInterface = baseInt
	fmt.Println(baseInterface.Show())

}

func testNil() {
	fmt.Println("test 接收者底层为nil值")
	var test *test_inter.BaseStruct
	var inter test_inter.BaseInter
	inter = test

	fmt.Println(inter.Show())

}

func testAnyInterface() {

	fmt.Println("test 空接口")

	var i interface{}

	i = 10
	fmt.Println(i)

	i = test_inter.BaseStruct{}

	// i 没实现Show方法
	//fmt.Println(i.Show())

	base := test_inter.BaseStruct{}

	fmt.Println(base.Show())
}

func testType(inVar interface{}) {
	switch t := inVar.(type) {
	case string:
		fmt.Println("is string", t)
	case int:
		fmt.Println("is int", t)
	default:
		fmt.Println("unknown")
	}
}

func testString() {
	strStru := test_inter.MyString{}
	fmt.Println("接收者：指向struct类型的指针实现了接口，并未在值类型的struct实现接口")
	fmt.Println(&strStru, strStru)
}

func testTag() {
	stu := new(test_inter.MyString)
	stu.SetName("io_demo test tag name")
	fmt.Println("测试tag")
	stu.ShowTag()
	fmt.Println(stu)
}
