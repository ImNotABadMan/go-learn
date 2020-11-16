package main

import (
	"demo/thread_demo/list"
	"fmt"
)

func main() {
	fmt.Println("紧急任务不可分割，优先级排列在多级队列中的一级紧急任务队列")
	fmt.Println("不紧急短任务，可中断，优先级排列和抢占，排在多级队列中的二级短任务队列")
	fmt.Println("不紧急耗时长任务，可中断，优先级排列和抢占，排在多级队列的三级长任务队列")
	highList := list.HighList{}
	shortList := list.ShortList{}
	longList := list.LongList{}
	listSlice := make([]interface{}, 3)
	listSlice[0] = highList
	listSlice[1] = shortList
	listSlice[2] = longList

	fmt.Println(highList)
	fmt.Println(shortList)
	fmt.Println(longList)

	fmt.Println("队列")
	fmt.Println(listSlice)
}
