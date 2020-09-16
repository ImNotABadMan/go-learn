package first_moudle

import (
	"errors"
	"fmt"
)

var gMsg string

type StuSrtuct struct {
	Id   uint64
	Name string
	sort int64
}

var Gs string

// 切片
var sliceArr []string

func Say(name string) string {
	message := fmt.Sprintf("name is : %s", name)
	return message
}

func SetError(name string) (string, error) {
	if name == "" {
		return fmt.Sprintf("%s", errors.New("Empty Name")), errors.New("Two Param Err")
	} else {
		gMsg = fmt.Sprintf("Not Empty SetError Name is %s", name)
	}

	return gMsg, nil
}

func StructShow() {
	type stu struct {
		id   int64
		name string
	}

	stus := stu{
		id:   10,
		name: "string",
	}

	fmt.Println(stus)
}

func ArrShow() {

	// 声明
	var arr [5]int64

	arr[0] = 1

	fmt.Println(arr)

	// 声明直接赋值
	arr1 := [5]string{"string", "s1", "s2"}

	fmt.Println(arr1)

	// 自动计算...
	arr2 := [...]byte{'2', '\x01'}

	fmt.Println(arr2)
}

func SetSlice(sliceStr []string) ([]string, error) {

	sliceArr = sliceStr

	sliceArr = append(sliceStr, "add 1")


	return sliceArr, nil
}

func GetSlice() {

	for _,v := range sliceArr {
		fmt.Println("Slice ", v)
	}

	fmt.Println(sliceArr)

}
