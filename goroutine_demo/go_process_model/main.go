package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	//test1()
	//test2()
	test3()
}

func test3()  {
	max := 16
	sLen := 12

	sliceInt := make([]string, 0)
	runtime.GOMAXPROCS(sLen)
	wait := sync.WaitGroup{}
	wait.Add(max)

	for index := 0; index < max; index++ {
		start := time.Now().UnixNano()
		//sliceInt = append(sliceInt, strconv.FormatInt(int64(index), 10) +
		//	" go start " + strconv.FormatInt(start, 10))
		go func(goInt []string, index int) {
			//start := time.Now().UnixNano()
			//sliceInt = append(sliceInt, strconv.FormatInt(int64(index), 10) +
			//	" go start " + strconv.FormatInt(start, 10))
			//fmt.Println(sliceInt)
			//h,m,s := time.Now().Clock()
			//fmt.Println(index, "go start" , h , ":", m, ":", s)
			time.Sleep(time.Millisecond)
			sliceInt = append(sliceInt, strconv.FormatInt(int64(index), 10) +
				" go end " + strconv.FormatInt(time.Now().UnixNano(), 10) + " 差：" +
				strconv.FormatInt(time.Now().UnixNano() - start, 10))


			//h,m,s = time.Now().Clock()
			//fmt.Println(index, "go end", time.Now().Unix(), h , ":", m, ":", s)

			wait.Done()
		}(sliceInt, index)
	}

	wait.Wait()
	fmt.Println("Finish")
	//fmt.Println(sliceInt)
	sort.Slice(sliceInt, func(i, j int) bool {
		return sliceInt[i] > sliceInt[j]
	})
	for _, item := range sliceInt{
		fmt.Println(item)
	}
}

func test2() {
	max := 16
	sLen := 12

	sliceInt := make([]int, sLen)

	runtime.GOMAXPROCS(sLen)

	wait := sync.WaitGroup{}
	wait.Add(max)
	ch := make(chan int)

	for index := 0; index < max; index++ {
		go func(goInt []int, index int) {
			res := <-ch
			time.Sleep(time.Millisecond * 2)
			wait.Done()
			fmt.Println(index, "index", res)
		}(sliceInt, index)
	}

	// 发送信号
	for index := 0; index < max; index++ {
		ch <- index
	}

	wait.Wait()
	fmt.Println("Finish")
	fmt.Println(sliceInt)
}

func test1() {
	max := 16
	sLen := 12

	sliceInt := make([]int, sLen)

	runtime.GOMAXPROCS(sLen)

	wait := sync.WaitGroup{}
	wait.Add(max)
	//mu := sync.Mutex{}
	for index := 0; index < max; index++ {
		go func(goInt []int, index int) {
			for i := 0; i < len(sliceInt); i++ {
				if sliceInt[i] == 0 {
					//mu.Lock()

					sliceInt[i]++
					time.Sleep(time.Millisecond)
					//mu.Unlock()
					fmt.Println(index)
					break
				}
			}

			time.Sleep(time.Millisecond * 2)
			wait.Done()
		}(sliceInt, index)
	}
	wait.Wait()
	fmt.Println("Finish")
	fmt.Println(sliceInt)
}
