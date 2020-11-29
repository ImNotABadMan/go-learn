package main

import (
	"errors"
	"fmt"
	"go-traning/error_demos/error_pointer"
	"go-traning/error_demos/error_value"
)

func main() {

	testErrorValue()
	testErrorPointer()

	testMetaError()



}

func testErrorValue() {
	err := error_value.New("test")
	err1 := error_value.New("test")

	metaError := errors.New("test")

	fmt.Println("err == err1", err == err1)
	fmt.Println("metaError", metaError == err)
}

func testErrorPointer() {
	fmt.Println("test error pointer")
	err := error_pointer.NewErrorPointer("test")
	err1 := error_pointer.NewErrorPointer("test")
	metaError := errors.New("test")

	fmt.Println("pointer err == err1", err == err1)
	fmt.Println("metaError err == err1", err == metaError)
}

func testMetaError() {
	err := errors.New("test")
	err1 := errors.New("test")

	fmt.Println("MetaError err == err1", err == err1)
}
