package main

import (
	"datatypes/interface_demo/test_inter"
)

func main() {
	testInterface()
}

func testInterface() {
	base := test_inter.BaseStruct{}
	base.Show()
}
