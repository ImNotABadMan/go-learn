package main

import (
	"fmt"
	"sort"
)

func main() {
	test()
	testExists()
	testSort()
	sortMap()

}

func init() {
	//声明
	var map1 map[int]string

	// 存储任意值
	var map2 map[int]interface{}

	type goods struct {
		id int64
	}

	var map3 map[goods]interface{}

	var map4 map[string]func() string

	map1 = map[int]string{}
	map1[1] = "test"

	good := goods{
		id: 2,
	}

	map3 = map[goods]interface{}{good: "test"}
	map3[goods{id: 2}] = "test"
	map3[goods{id: 1}] = 1234

	map4 = map[string]func() string{}
	map4["test"] = func() string {
		fmt.Println("fmt print test mp4")
		return "test mp4"
	}

	var map5 = make(map[int]string, 50)

	map5[2] = "ttt"

	// 切片值类型
	map6 := map[string]*[]int{}
	map6["map6"] = &[]int{1, 2, 3}

	map7 := map[string]*[]int{}
	map7["map7"] = &[]int{1, 2, 3}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println(map4["test"]())
	fmt.Println(map5)
	fmt.Println(map7)
	fmt.Println(*map7["map7"])

	fmt.Println(map6)
	fmt.Println(*map6["map6"])

	fmt.Println("init func content ---------- ")
}

func test() {
	//如果你错误的使用 new () 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	var test = new(map[int]string)

	fmt.Println(test)

	*test = map[int]string{1: "test"}
	(*test)[2] = "test2"
	fmt.Println(test)

	fmt.Println(*test)
}

func testExists() {
	fmt.Println("testExists -------------")
	map1 := make(map[int]string)
	map1[1] = "test"
	value, keyOk := map1[2]
	value1, key1Ok := map1[1]
	fmt.Println("keyOk:", keyOk, "value:", value)
	fmt.Println("key1Ok:", key1Ok, "value1:", value1)

}

func testSort() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

func sortMap() {

	fmt.Println("sortMap --------- ")

	map1 := map[string]int{"test2": 2, "atest5": 5, "test1": 1}
	slice1 := []string{}

	for key := range map1 {
		slice1 = append(slice1, key)
	}

	fmt.Println(map1)
	fmt.Println(slice1)

	sort.Strings(slice1)

	fmt.Println(slice1)

	for _, value := range slice1 {
		fmt.Print("\t", map1[value])
	}

	fmt.Println()
}
