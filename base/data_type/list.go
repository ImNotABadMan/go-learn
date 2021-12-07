// 声明一个main包（包是对功能进行分组的一种方式，它由同一目录中的所有文件组成）。
package main

import (
	inline_package "base/data_type/inline-package"
	"fmt"
)

func main() {
	fmt.Println("second mod")

	str := inline_package.ShowInline()

	fmt.Println(str)
}
