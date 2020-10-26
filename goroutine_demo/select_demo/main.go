package main

import "fmt"

func selectFunc(strCh, quit chan string) {
	select {
	case str := <-strCh:
		fmt.Println(str)
	case <-quit:
		fmt.Println("quit...")
	default:
		fmt.Println("default")
	}
}

func main() {
	strCh := make(chan string, 1)
	quit := make(chan string)
	selectFunc(strCh, quit)
	strCh <- "test"
	selectFunc(strCh, quit)
}
