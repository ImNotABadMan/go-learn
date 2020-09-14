package main

import (
	"fmt"
	"hello"
	"log"
	"math"
	"os"
	"runtime"
	"strings"
	"time"
)

const c1 = 0.111111111111111111111111

// 全局变量
var global string

type inStruct struct {
	id   int
	name string
}

// 闭包调试
var where = func() {
	caller, file, line, ok := runtime.Caller(1)
	log.Println(caller, file, line, ok)
}

var calTime = func() func() {
	var start = time.Now()
	var end time.Time
	return func() {
		if end.IsZero() {
			end = time.Now()
		} else {
			fmt.Printf("Time: %s\n", end.Sub(start).String())
			start, end = time.Now(), time.Now()
		}
	}
}

// test
func main() {
	var str string
	if len(os.Args) > 1 {
		str = os.Args[1]
	}
	fmt.Println(str, hello.Reserve(str))

	// 指针
	a := "testaaaaa"
	var b *string = &a
	fmt.Println("a: ", a, " &b: ", b, " b: ", **(&b))
	*b = "test"
	fmt.Println("a: ", a, " &b: ", b, " b: ", *b)

	fmt.Println(strings.HasPrefix(a, "tes"))
	fmt.Println("函数")
	t := 2

	st := test(1, &t)
	fmt.Println("引用传递：", t)
	fmt.Println(st)
	x1, x2 := mingmingfahui(5)
	fmt.Println(x1, x2)

	// 参数变长函数
	slice := []int{1, 2, 3}
	test1(slice...)

	ccbqd(1, "s", 23, 0.011)

	// 回调函数
	varHuidDiao := huiDiao(func(i int) int32 {
		pow := math.Pow(float64(i), 4)
		return int32(pow)
	})

	fmt.Println(varHuidDiao)

	// 匿名函数
	niMing("NiMing")

	// defer
	fmt.Println(testDefer())

	// 闭包作用域，以及变量持久化状态
	biBao := testBiBao()
	fmt.Println(biBao(1))
	fmt.Println(biBao(5))
	fmt.Println(biBao(10))

	// 工厂模式
	fmt.Println("工厂模式")
	fac := factory()
	fmt.Println(fac())
	fmt.Println(fac())
	fmt.Println(fac())

	calTimeFunc := calTime()
	calTimeFunc()
	where()
	calTimeFunc()

}

// 函数
func test(a int, b *int) inStruct {

	var aa = b
	bb := &a

	fmt.Println(a, b)
	fmt.Println(&a, *b)
	fmt.Println(aa, bb)
	fmt.Println(*aa, *bb)

	*b = 10

	return inStruct{
		*b,
		string(a),
	}
}

// 参数变长函数
func test1(...int) string {
	return string(global)
}

func mingmingfahui(x int) (x1 int, x2 int) {
	x1 = x * x
	x2 = x + x
	return
}

// 参数类型不确定的变长函数
func ccbqd(values ...interface{}) {
	// 循环遍历参数
	for _, v := range values {
		// 判断类型
		switch v.(type) {
		case int:
			fmt.Println(v)
		default:
			fmt.Println("default: ", v)
		}
	}
}

// 回调函数
func huiDiao(f func(int) int32) int32 {
	return f(2)
}

// 匿名函数
func niMing(str string) {
	varNiming := func(s string) string {
		return strings.Join([]string{"test", s}, "-")
	}(str)

	fmt.Println(varNiming)
}

// defer 结构
func testDefer() (int, error) {
	fmt.Println("test Defer函数")
	for index := 0; index < 10; index++ {
		defer fmt.Println(index)
		defer niMing(strings.Join([]string{fmt.Sprintf("%d", index), "q"}, "."))
	}
	index1 := 1
	defer fmt.Println("func")

	// defer 在return后执行
	defer func(i *int) {
		fmt.Println("defer func()")
		*i++
	}(&index1)

	return fmt.Println(index1)
}

// 函数返回闭包，函数内的变量被固定了
func testBiBao() func(int) int {
	var x int
	return func(index int) int {
		x = index + x
		return x
	}
}

// 工厂模式
func factory() func() inStruct {
	var fac inStruct
	return func() inStruct {
		if fac.id == 0 {
			fac = inStruct{
				id:   1,
				name: "Connect",
			}
		}
		return fac
	}
}
