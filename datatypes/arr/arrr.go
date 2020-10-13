package main

import "fmt"

func main() {
	fmt.Println("test")
	slice := testSlice()
	fmt.Println("slice: ", slice)
}

// 切片
func testSlice() []int {
	return []int{1, 2, 3}
}
