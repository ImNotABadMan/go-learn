package main

import "fmt"

type TZ int

const (
	UTC = 1
	ETC = 2
)

func main() {
	fmt.Println("test")
	slice := testSlice()
	fmt.Println("slice: ", slice)

	arr := []int{UTC, ETC}
	arr1 := []TZ{3, 4}

	fmt.Println(arr, arr1)

}

// 切片
func testSlice() []int {
	return []int{1, 2, 3}
}
