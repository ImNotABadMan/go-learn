package test_inter

import "fmt"

type BaseInter interface {
	Show() string
}

type BaseStruct struct {
	name string
}

func (base *BaseStruct) Show() string {
	if base == nil {
		return "<nil>"
	}
	fmt.Println("test base interface", base.name)
	return "test"
}

type MyInt int

func (this MyInt) Show() string {
	fmt.Println("test MyInt Interface", this)
	return "int test"
}
