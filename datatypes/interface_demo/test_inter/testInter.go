package test_inter

import "fmt"

type baseInter interface {
	Show() string
}

type BaseStruct struct {
	name string
}

func (base *BaseStruct) Show() string {
	fmt.Println("test base interface")
	return "test"
}
