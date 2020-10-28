package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var p unsafe.Pointer
	p = nil
	fmt.Println(p)

	testStr := struct {
		name string
	}{
		name: "test",
	}

	(&testStr).name = "testName"

	// 指针的十进制表示
	var p2 uintptr

	ref := reflect.ValueOf(&testStr)
	p2 = ref.Pointer()

	p3 := unsafe.Pointer(&testStr)

	fmt.Println(ref, p2, p3, uintptr(p3))

	// 用内存地址偏移量得出结构体中的字段地址，然后取值，利用指针原理
	testP := &testStr
	namePOffset := unsafe.Offsetof(testP.name)
	nameS := uintptr(p2) + namePOffset

	fmt.Println("name value")
	fmt.Println(namePOffset, nameS, (*string)(unsafe.Pointer(&nameS)), unsafe.Pointer(&nameS))

	p2 = 123456
	fmt.Println(p2)

	//
	const c1 = 123

	addend := 1
	//(5+5)++
	addend++
	for index := 0; index < 5; index++ {
		addend++
	}

}
