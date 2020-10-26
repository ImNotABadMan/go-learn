package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("cores cal")
	fmt.Println("利用go协程进行核和运算")

	params := flag.Int("slice-len", 100, "切片的长度")
	//part := flag.Int("parts", 5, "切片的分片数")
	cores := flag.Int("cores", 2, "计算核心数")
	flag.Parse()

	runtime.GOMAXPROCS(*cores)

	var (
		sliceGo, sliceSync []int
		resCh              chan int64
		count              int
	)

	sliceSync, sliceGo, resCh = make([]int, 0), make([]int, 0), make(chan int64)

	for i := 0; i < *params; i++ {
		sliceGo = append(sliceGo, i)
		sliceSync = append(sliceSync, i)
	}

	var start = time.Now()

	//go goCal(*part, sliceGo, resCh)
	go syncCal(sliceSync, resCh)

	for {
		tmp := <-resCh
		fmt.Println("end: ", count, time.Now())
		fmt.Println("运行了: ", count, time.Now().Sub(start).Microseconds(), "微秒，", time.Now().Sub(start).Seconds(), "秒")
		fmt.Println("resCh", count, tmp)
		fmt.Println()
		count++
		if count == 2 {
			break
		}
	}

	fmt.Println("test end .............")
}

func goCal(parts int, calSlice []int, outCh chan int64) {
	fmt.Println("协程分隔计算")
	fmt.Println("goCal start at:", time.Now())

	var sumAll int64
	// 切片的长度
	length := len(calSlice)
	// 最后一个剩余的余数
	var lastExtra = length % parts
	// 每部分的切片长度
	var partLen = int(length / parts)
	// 管道：用于协程直接的通信，功能：存储计算的结果
	resCh := make(chan int64)
	// 函数变量
	type cal func(partSlice []int) int64
	// 计算函数
	var calPartSum cal = func(partSlice []int) int64 {
		var sum int64
		for _, val := range partSlice {
			sum += int64(val)
		}
		resCh <- sum

		return sum
	}

	for step := 0; step < parts; step++ {
		start, end := step*partLen, (step+1)*partLen
		if step == parts-1 {
			end += lastExtra
		}
		//fmt.Println("go ", start, ":", end)
		go calPartSum(calSlice[start:end])
	}

	fmt.Println("go length:", length)
	fmt.Println("go lastExtra:", lastExtra)
	fmt.Println("go partLen:", partLen)

	var index int

	for partSum := range resCh {
		//fmt.Println("part sum:", partSum)

		sumAll += partSum
		if index++; index >= parts {
			break
		}
	}

	outCh <- sumAll

	fmt.Println("go sum:", sumAll)
}

func syncCal(calSlice []int, outCh chan int64) int64 {
	fmt.Println("串行进行")
	fmt.Println("syncCal start at:", time.Now())
	var sum int64
	for _, value := range calSlice {
		sum += int64(value)
	}
	outCh <- sum

	fmt.Println("sync sum:", sum)

	return sum
}
