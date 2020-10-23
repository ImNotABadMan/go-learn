package test_inter

import (
	"fmt"
	"reflect"
)

type MyString struct {
	name string "test name"
}

func (this *MyString) String() string {
	return this.name + ",haha haha"
}

func (this *MyString) ShowTag() {
	fmt.Println("使用带标签的struct, tag 只能在包里使用")
	tType := reflect.TypeOf(*this)
	for i := 0; i < 1; i++ {
		field := tType.Field(i)
		fmt.Println("Tag:", field.Tag)
	}
}

func (this *MyString) SetName(str string) {
	this.name = str
}
