package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	testChange()
}

func testChange() {
	str := "origin string"
	//在 Go 语言中，可供我们选择的同步工具并不少。
	//其中，最重要且最常用的同步工具当属互斥量（mutual exclusion，简称 mutex）。
	//sync包中的Mutex就是与其对应的类型，该类型的值可以被称为互斥量或者互斥锁。
	// 变成串行执行
	lock := sync.Mutex{}
	go change(&str, "first string", &lock, func(inStr string) {
		//time.Sleep(time.Millisecond)
		if inStr == "first string" {
			fmt.Println("is first string")
		}
	})
	go change(&str, "second string", &lock, func(string2 string) {})
	fmt.Println(str)
	time.Sleep(time.Millisecond * 2)
	fmt.Println(str)
}

func change(str *string, newStr string, lock *sync.Mutex, f func(string2 string)) {
	lock.Lock()
	*str = newStr
	fmt.Println("change:", *str)
	f(*str)
	lock.Unlock()
}

//
//// 接口类型 是由一组方法签名定义的集合。
////
////接口类型的变量可以保存任何实现了这些方法的值。
//func change(str *string, newStr string, lock *sync.Mutex, v interface{})  {
//	lock.Lock()
//	*str = newStr
//	fmt.Println("change:", *str)
//	switch t := v.(type) {
//	case func():
//		t()
//	default:
//		fmt.Println("type is:", t)
//	}
//	lock.Unlock()
//}
